#!/usr/bin/env python3

template_dir = "D:/Tillo/Documents/hw_altium/trunk/Templates"
templates = ["twolayer", "fourlayer", "sixlayer"]

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
logging.setLevel("INFO")

def get_templates():
    logging.warning("Copying template sources...")
    for t in templates:
        logging.warning("\tCopying {}...".format(t))
        shutil.copytree(os.path.join(template_dir,t),os.path.join("templates_src",t))
        for dir in ["__Previews",".svn", "History", "Project Logs for "+t,"Project Outputs for "+t]: # DIRS to clean up
            try:
                logging.info("\t\tdeleting {}".format(dir))
                shutil.rmtree(os.path.join("templates_src",t,dir))
            except:
                pass

def compress_templates():
    logging.warning("Compressing templates...")
    for dir in os.listdir("templates_src"):
        try:
            hist = os.path.join(os.getcwd(),"templates_src",dir,"History")
            try:
                shutil.rmtree(hist,ignore_errors=True)
            except FileNotFoundError:
                logging.info("\t{} not found".format(hist))

            hist = os.path.join(os.getcwd(),"templates_src",dir,"__Previews")
            try:
                shutil.rmtree(hist,ignore_errors=True)
            except FileNotFoundError:
                logging.info("\t{} not found".format(hist))

            logging.info("\tzipping {}".format(dir))
            shutil.make_archive("templates/"+dir, 'zip', "templates_src/"+dir)
        except NotADirectoryError:
            pass

def cleanup():
    logging.warning("Deleting template sources...")
    shutil.rmtree("templates_src")

    logging.warning("Deleting compressed templates...")
    shutil.rmtree("templates")


get_templates()
compress_templates()

logging.warning("Installing using pakr...")
subprocess.call("packr2 install")

cleanup()