Here's a clean solution for archiving live streaming data with a 10MB file rotation:Here's a clean solution for archiving live streaming data with a 10MB file rotation:Here's how this works and a ready-to-use Python implementation:

**The core idea** — you continuously write to a single active file, check its size after each write, and when it hits 10 MB you archive it (rename/compress), then start a fresh file.

```python
import os
import gzip
import shutil
from datetime import datetime

MAX_SIZE = 10 * 1024 * 1024  # 10 MB
ACTIVE_FILE = "stream_data.log"
ARCHIVE_DIR = "archives"

os.makedirs(ARCHIVE_DIR, exist_ok=True)

def archive_and_rotate(filepath):
    """Compress and move the full file, return a fresh one."""
    timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
    archive_name = os.path.join(ARCHIVE_DIR, f"stream_{timestamp}.log.gz")
    
    with open(filepath, "rb") as f_in:
        with gzip.open(archive_name, "wb") as f_out:
            shutil.copyfileobj(f_in, f_out)
    
    # Truncate the active file (keeps the handle open safely)
    open(filepath, "w").close()
    print(f"Archived → {archive_name}")

def write_stream(data: bytes):
    """Write data and rotate if over 10 MB."""
    with open(ACTIVE_FILE, "ab") as f:
        f.write(data)
    
    if os.path.getsize(ACTIVE_FILE) >= MAX_SIZE:
        archive_and_rotate(ACTIVE_FILE)

# --- Simulate a live stream ---
import time
chunk = b"x" * 1024  # 1 KB per chunk

while True:
    write_stream(chunk)
    time.sleep(0.01)  # adjust to your feed rate
```

**Key design decisions:**

- **Truncate, don't delete** — `open(filepath, "w").close()` empties the file in-place. This is safer than deleting and recreating if another process holds the file open.
- **Compress on archive** — `gzip` cuts storage by 60–90% on text/log data. If you need random access later, skip it and just rename.
- **Timestamped archives** — `stream_20240528_143022.log.gz` makes ordering and cleanup easy.
- **Thread-safe variant** — if your stream writes from multiple threads, wrap the size check + archive in a `threading.Lock()` to avoid double-archiving.

**For very high-throughput streams** (e.g. GB/s), replace the per-write `os.path.getsize()` call with an in-memory counter to avoid the syscall overhead:

```python
_written = 0

def write_stream(data: bytes):
    global _written
    with open(ACTIVE_FILE, "ab") as f:
        f.write(data)
    _written += len(data)
    if _written >= MAX_SIZE:
        archive_and_rotate(ACTIVE_FILE)
        _written = 0
```

Click any box in the diagram above to go deeper on any step.