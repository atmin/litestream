package internal

import "testing"

// TestFsyncDir is a portability regression test. The POSIX implementation
// (os.Open followed by Sync) cannot work on Windows — a directory handle there
// is not open for writing, so the sync fails with "Access is denied" — and a
// restore is where that surfaced: litestream fsyncs the output directory after
// renaming the restored database into place, so every restore on Windows failed
// at the last step. Asserting only that the call succeeds is the whole point,
// since that is precisely what did not hold.
//
// The two arms are deliberately not equivalent for a missing directory: the
// POSIX one reports the error from opening it, the Windows one does nothing at
// all. Every caller runs this immediately after renaming a file *into* the
// directory, so it exists by construction.
func TestFsyncDir(t *testing.T) {
	if err := FsyncDir(t.TempDir()); err != nil {
		t.Fatalf("FsyncDir on an existing directory: %v", err)
	}
}
