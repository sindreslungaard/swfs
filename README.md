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

#### Extract
The extract tool dumps all png and binary files of the swf files in a directory.

| Argument  | Explanation |
| ------------- | ------------- |
| input       | Path to the directory where the swf files are located |
| output         | Path to the directory you want the extracted files to be placed  |
| [workers] | Size of worker pool. Higher number will increase the concurrent use of swfdump and swfextract. Default is 2.

```bash
./extract -input /swfdir -output /outputdir -workers 2
```