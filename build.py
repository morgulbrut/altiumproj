#!/usr/bin/env python3


import os
import colorlog
import shutil
import subprocess

handler = colorlog.StreamHandler()
handler.setFormatter(colorlog.ColoredFormatter(
    "%(log_color)s%(message)s",
    datefmt=None,
    reset=True,
    log_colors={
        'DEBUG':    'cyan',
        'INFO':     'green',
        'WARNING':  'yellow',
        'ERROR':    'red',
        'CRITICAL': 'red,bg_white',
    },
    secondary_log_colors={},
    style='%'
))

logging = colorlog.getLogger('-')
logging.addHandler(handler)


def compress_templates():
    logging.warning("Deleting old templates...")
    for file in os.listdir("templates"):
        logging.info("\tdeleting {}".format(file))
        os.remove("templates/"+file)

    logging.warning("Compressing templates...")
    for dir in os.listdir("templates_src"):
        logging.info("\tzipping {}".format(dir))
        shutil.make_archive("templates/"+dir, 'zip', "templates_src/"+dir)


compress_templates()
logging.warning("Installing using pakr...")
subprocess.call("packr2 install")
