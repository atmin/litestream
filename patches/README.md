# Carried patches

This directory is the recovery copy of the `s3lite` branch: `git format-patch`
output for every commit the branch carries on top of its upstream base, in apply
order. If the branch is ever lost or force-pushed wrong, rebuild it from here:

```bash
git checkout -b s3lite <upstream-release-tag>
git am --3way patches/*.patch
```

The authoritative ledger — what each patch is for, and whether it has an upstream
PR that would let it be dropped — lives in
[s3lite's LITESTREAM-FORK.md](https://github.com/atmin/s3lite/blob/master/LITESTREAM-FORK.md).

Regenerate after changing the branch:

```bash
rm -f patches/*.patch
git format-patch <upstream-release-tag>..s3lite -o patches \
  --no-signature -N --zero-commit -- . ':(exclude)patches'
```

Two flags keep the series a fixpoint rather than churn: `patches/` excludes itself
(else every refresh rewrites the patch that carries the previous refresh), and
`--zero-commit` blanks the SHA in each patch header (else every rebase rewrites
every patch even when no code changed).

`.github/workflows/sync-upstream.yml` does this automatically on every clean
rebase onto a new upstream release.
