## swfs
swfs is a set of cli tools to work with swf files.

### 1. Download swfs
Download the [latest release](https://github.com/sindreslungaard/swfs/releases/latest) of swfs for your operating system of choice.


### 2. Install `swftools` (http://www.swftools.org/)

swfs uses swftools under the hood, make sure you have it installed on your system and added to your PATH.

|[swftools windows](http://www.swftools.org/swftools-2013-04-09-1007.exe)|[swftools linux](http://www.swftools.org/swftools-2013-04-09-1007.tar.gz)|
| ------------- | ------------- |

`swfdump` and `swfextract` should be runnable from your commandline.

### 3. Usage

Explore the different tools below by expanding the details.

##### Extracter `./swfs [options] extract`
<details>
The extract tool dumps all png and binary files of the swf files in a directory.

| Argument  | Explanation |
| ------------- | ------------- |
| input       | Path to the directory where the swf files are located |
| output         | Path to the directory you want the extracted files to be placed  |
| [workers] | Size of worker pool. Higher number will increase the concurrent use of swfdump and swfextract. Default is 2.

```bash
./swfs -input /swfdir -output /extracted -workers 2 extract
```
</details>

##### Bundler `./swfs [options] bundle`
<details>
The bundle tool replaces the extracted folders into individual  `.asset` files.

| Argument  | Explanation |
| ------------- | ------------- |
| input         | Path to the directory where the extracted files are located |
| [workers] | Size of worker pool. Higher number will increase the concurrency of the program. Default is 5.|

```bash
./swfs -input ./extracted -workers 5 bundle
```

The `.asset` file structure:
```bash
## Content of .asset files are structured as key value pairs
## separated by "=\n". Multiple assets are separated with 
## double newline "\n\n". First line is always the format version.

version=
1


some_extracted_image.png=
�PNG...


some_extracted_binary.bin=
<xml>
    ....
</xml>
```

To parse `.asset` files, split on `\n\n`, then loop through each part. Split each part on `=\n` to get [0]filename and [1]filedata. 
</details>