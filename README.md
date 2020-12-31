## swfs
swfs is a set of cli tools towork with swf files.

### 1. Download the latest release of swfs
....

### 2. Install `swftools` (http://www.swftools.org/)

For the asset dumping & extraction process to work correctly, `swftools` needs to be installed in your system.
Download `swftools` [here](http://www.swftools.org/download.html) and install it.

| OS  | Download link |
| ------------- | ------------- |
| Windows       | http://www.swftools.org/swftools-2013-04-09-1007.exe  |
| Linux         | http://www.swftools.org/swftools-2013-04-09-1007.tar.gz  |


After installation, add the installation directory of `swftools` to your systems `PATH` variable.
In the end, the commands `swfdump` and `swfextract` should be callable from the command line.

### 3. Usage

Explore the different tools below by expanding the details.

##### Extracter
<details>
The extract tool dumps all png and binary files of the swf files in a directory.

| Argument  | Explanation |
| ------------- | ------------- |
| input       | Path to the directory where the swf files are located |
| output         | Path to the directory you want the extracted files to be placed  |
| [workers] | Size of worker pool. Higher number will increase the concurrent use of swfdump and swfextract. Default is 2.

```bash
./extract -input /swfdir -output /extracted -workers 2
```
</details>

##### Bundler
<details>
The bundle tool replaces the extracted folders into individual  `.asset` files.

| Argument  | Explanation |
| ------------- | ------------- |
| input         | Path to the directory where the extracted files are located |
| [workers] | Size of worker pool. Higher number will increase the concurrency of the program. Default is 5.|

```bash
./bundle -input ./extracted -workers 5
```

The `.asset` file structure:
```bash
## Content of .asset files are structured as key value pairs
## separated by "=\n". Multiple assets are separated with 
## double newline "\n\n"

version=1 # Format version used, always first line


some_extracted_image.png=
�PNG...


some_extracted_binary.bin=
<xml>
    ....
</xml>
```

To parse `.asset` files, split on `\n\n`, then loop through each part. Split each part on `=\n` to get [0]filename and [1]filedata. 
</details>