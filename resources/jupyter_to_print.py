# # TODO: Put this in a docker container to manage dependencies
#! /usr/bin/env python

from os.path import splitext, extsep, dirname, abspath, basename, exists
from os import curdir, chdir, remove
import argparse
import sys

import time
from glob import glob
from tqdm import tqdm

from traitlets.config import Config
from nbconvert.preprocessors import ExecutePreprocessor
import nbformat
from nbconvert import HTMLExporter


parser = argparse.ArgumentParser(
    description='Reportify: Jupyter notebook to report-formatted HTML. '
                'Optimized for pasting into a Google doc')
parser.add_argument('in_file_path', metavar='F', type=str, nargs='+',
                    help='The path to the notebook file to be converted')
parser.add_argument('-pm', '--with-pymarkdown', default=False,
                    action='store_true',
                    help='Enable PyMarkdown processing')
parser.add_argument('--no-embed-images', default=False,
                    action='store_true',
                    help='Disable HTML image embedding')
parser.add_argument("-o", "--output",
                    metavar='O',
                    type=str,
                    default=dirname(abspath(__file__)))
args = parser.parse_args()

this_file_dir_path = dirname(abspath(__file__))

c = Config()

if args.with_pymarkdown:
    try:
        # Add preprocessor to do markdown-python rendering.
        from jupyter_contrib_nbextensions.nbconvert_support import pre_pymarkdown
    except ImportError:
        print('Cannot enable PyMarkdown support as it is not installed')
    else:
        c.HTMLExporter.preprocessors.append(
            pre_pymarkdown.PyMarkdownPreprocessor)
else:
    print('Not using PyMarkdown')

if not args.no_embed_images:
    from embed_html import EmbedHTMLExporter
    Exporter = EmbedHTMLExporter
else:
    Exporter = HTMLExporter

# Template lives next to this build file, so add that to search path.
c.HTMLExporter.template_path.append(this_file_dir_path)
# Tell it to use our custom HTML template.
c.HTMLExporter.template_file = 'no_code'

exporter = Exporter(config=c)

for idx, file_path in enumerate(args.in_file_path):
    sys.stdout.write(file_path)
    # file_path = abspath(file_path).replace(' ', '\ ')

    root = dirname(file_path)
    if root == '':
        root = '.'
    title = basename(file_path).split('.ipynb')[0]

#     # Remove pre-existing versions of this post if it's there
    # existing = glob('{}{}.html'.format(html_path, title))
    # if len(existing) > 1:
    #     raise ValueError('There should be at most 1 file with this title')
    # elif len(existing) == 1:
    #     # Keep the date info for the post
    #     existing = existing[0]
    #     date_info = basename(existing).split('-')[:3]
    #     year, mo, day = [int(ii) for ii in date_info]
    #     remove(existing)
    # else:
    #     # Create new date info for the post
    now = time.localtime()
    mo = now.tm_mon
    day = now.tm_mday
    year = now.tm_year

    # Make output file name
    # out_file_name = '{}{}html'.format(splitext(file_path)[0], extsep)
    out_file_name = '{}-{}-{}-{}.html'.format(year, mo, day, title)
    file_dir = dirname(abspath(file_path))
    chdir(file_dir)  # nbconvert is weird
    out_file = "{}/{}".format(args.output, out_file_name)

    with open(file_path, "r") as read_fp:
        nb = nbformat.read(read_fp, as_version=nbformat.NO_CONVERT)
        ep = ExecutePreprocessor(timeout=600)
        ep.preprocess(nb, {'metadata': {'path': file_dir}})

        # CLEAN UP CLEAN UP
        new_cells = []
        for ii, cell in enumerate(nb.cells):

            # Add non-code cells
            if cell['cell_type'] != 'code':
                new_cells.append(cell)
                continue

            # Skip empty cells
            if len(cell['source']) == 0:
                continue

            # Clean up outputs
            outputs = []
            for output in cell['outputs']:
                # Remove stderrs
                if 'name' in list(output.keys()):
                    if output['name'] != 'stderr':
                        continue

                # Check for object output
                if 'data' in list(output.keys()):
                    if 'text/plain' in output['data'].keys():
                        if output['data']['text/plain'].startswith('<'):
                            _ = output['data'].pop('text/plain')

                outputs.append(output)

            cell['outputs'] = outputs
            cell['execution_count'] = None

            new_cells.append(cell)

        nb['cells'] = new_cells

        resources = {'config_dir': this_file_dir_path}
        # Render the notebook. The resources, we will not use.
        (body, resources) = exporter.from_notebook_node(nb, resources=resources)

        try:
            print("Writing to {}".format(out_file))
            with open(out_file, mode='w') as f:
                f.write(body)
        except Exception as error:
            print("There was an error: {}".format(error))

    chdir(this_file_dir_path)
