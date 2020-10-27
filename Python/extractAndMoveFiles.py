
import os
import shutil
from os.path import isfile, join
from shutil import copy

f = []
srcPath = "/Users/varunvijayakumar/Downloads/srcDevices/devices"
srcFiles = [f for f in os.listdir(srcPath) if isfile(join(srcPath, f))]
#print(onlyfiles)
#print(len(srcFiles))
for item in srcFiles :
    item = item.split(".yaml")[0]
    #print(item)

dstPath = "/Users/varunvijayakumar/Desktop/FILES/NESTLE/UTILITIES/GoLang/yamlToJSON/devices"
for file in srcFiles:
    item = file.split(".yaml")[0]
    finalDstPath  = dstPath + "/" + item + "/"
    #print(finalDstPath)
    #print(file)
    dst = finalDstPath+"ports.yaml"
    src = srcPath+"/"+file
    #print("source " + src)
    #print("destination " + dst)
    #copy(src, dst)


def createDirectoryIfNotThere(dest, directory) :
    parent_dir = "D:/Pycharm projects/"

    # Path
    path = os.path.join(dest, directory)
    print(dest, directory)
    print(path)
    # Create the directory
    # 'GeeksForGeeks' in
    # '/home / User / Documents'
    os.mkdir(path)
    print("Directory '% s' created" % path)


movdir = r"/Users/varunvijayakumar/Desktop/FILES/NESTLE/UTILITIES/GoLang/yamlToJSON/devices"
basedir = r"/Users/varunvijayakumar/Downloads/srcDevices/devices"
# Walk through all files in the directory that contains the files to copy
for root, dirs, files in os.walk(basedir):
    for filename in files:
        # I use absolute path, case you want to move several dirs.
        item = filename.split(".yaml")[0]

        old_name = os.path.join( os.path.abspath(root), filename )

        # Separate base from extension
        base, extension = os.path.splitext(filename)
        print("base : " + base + " extension : " + extension)
        # Initial new name
        #base = base + "/ports.yaml"
        if ".yaml" not in extension :
            print("not yaml, ext"+ extension)
            continue
        new_name = os.path.join(movdir, base, "ports.yaml")

        print("fileName : "+ item +"old_name " + old_name + " new_name : "+ new_name)

        # create directory if it doesn't exist
        if not os.path.exists(os.path.join(movdir, base)):
            createDirectoryIfNotThere(movdir, base)

        #shutil.copy(old_name, new_name)

        # If folder basedir/base does not exist... You don't want to create it?
        if not os.path.exists(os.path.join(basedir, base)):
            print(os.path.join(basedir,base), "not found, creating ")
            shutil.copy(old_name, new_name)
            continue    # Next filename
        elif not os.path.exists(new_name):  # folder exists, file does not
            shutil.copy(old_name, new_name)
        else:  # folder exists, file exists as well
            ii = 1
            while True:
                new_name = os.path.join(basedir,base, base + "_" + str(ii) + extension)
                if not os.path.exists(new_name):
                   shutil.copy(old_name, new_name)
                   print("Copied", old_name, "as", new_name)
                   break
                ii += 1


