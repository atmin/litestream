//go:build windows
// +build windows

package internal

import (
	"os"
)

// Fileinfo returns syscall fields from a FileInfo object.
func Fileinfo(fi os.FileInfo) (uid, gid int) {
	return -1, -1
}

// FsyncDir is a no-op on Windows, because the operation does not exist there:
// flushing buffers needs a handle opened for writing, and a directory handle
// (which itself requires FILE_FLAG_BACKUP_SEMANTICS) is not one — so the unix
// implementation fails with "Access is denied" rather than syncing anything.
//
// Skipping it is safe rather than merely unavoidable. What fsync-on-directory
// buys on POSIX is that the rename entry reaches disk; on NTFS a rename is a
// journaled metadata transaction, so the directory entry is recovered by the
// log. The ordering that matters — the file's own contents landing before the
// rename that publishes them — comes from fsyncing the *file*, which every
// caller here already does.
func FsyncDir(string) error { return nil }

// fixRootDirectory is copied from the standard library for use with mkdirAll()
func fixRootDirectory(p string) string {
	if len(p) == len(`\\?\c:`) {
		if os.IsPathSeparator(p[0]) && os.IsPathSeparator(p[1]) && p[2] == '?' && os.IsPathSeparator(p[3]) && p[5] == ':' {
			return p + `\`
		}
	}
	return p
}
