#!/usr/bin/env python

import argparse
import json
import logging
import os.path
import subprocess

DEFAULT_OUTFILE = "result.json"

here = os.path.dirname(os.path.realpath(__file__))
ghlicense_bin = os.path.realpath(f'{here}/../ghlicense')

logging.basicConfig(level=logging.INFO, format='%(asctime)s %(message)s')
LOG = logging.getLogger(__name__)


# map of repo_url -> SPDX license
succeeded = {}
# map of repo_url -> error output
failed = {}

def store(file_path, result):
    # dump result to file
    with open(file_path, "w") as f:
        json.dump(result, f, indent=2)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("repo_list", metavar="repo-list", help="A file with GitHub repo URLs")
    parser.add_argument("--outfile", metavar="PATH", default=DEFAULT_OUTFILE, help="Output file.")
    args = parser.parse_args()

    result = {
        'ok': succeeded,
        'errors': failed
    }
    with open(args.repo_list, "r") as f:
        linum = 0
        for line in f:
            linum += 1
            repo_url = line.strip()
            if repo_url.startswith("#"):
                continue

            LOG.info("%d: %s", linum, repo_url)
            cmd = f'{ghlicense_bin} {repo_url}'
            p = subprocess.run(cmd, shell=True, capture_output=True)
            if p.returncode == 0:
                succeeded[repo_url] = p.stdout.decode('utf-8').strip()
            else:
                failed[repo_url] = p.stderr.decode('utf-8').strip()

            store(args.outfile, result)

    # dump to stdout
    print(json.dumps(result, indent=2))
