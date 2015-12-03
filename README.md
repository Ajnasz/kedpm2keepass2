# Kedpm to KeePass2

Export passwords from kedpm to import them into keepass2 using the Generic CSV Importer format.

```sh
$ echo export | kedpm -c > export.txt
```

```sh
$ cat export.txt | kedpm2keepass2 > import.csv
```

or in one line:

```
$ echo export | kedpm -c | kedpm2keepass2 > import.csv
```

Then open keepass2, open the File menu, choose _Import_. From the dialog choose _General CSV Importer_ item, and browse the _import.csv_ file. Press Ok, on the next screen choose the _Structure_ tab, remove the _Group_ item from the _Semantics_ section. Add a new Group field but select `/` as separator, move it to the top of the list, then press Next and if it looks good press Finish.
Finally remove the exported files:

```
$ rm import.csv export.txt
```
