# GoKOM

A simple ElSword KOM Archive extractor written in Go because experimenting with it.

The archiver is capable of extracting all the KOMs, simply create the directories '<KOM name>_out' and place it inside the ElSword data folder, from there just execute it.

It will extract all the files, but encryption for Algorithm 3 have been removed.
It will extract and decompress the files, tho.

# Warning!

Also DDS/TGA files are now being encrypted despite the algorithm field value, KOG's decision on removing modding communities.

# License

Do whatever you want, have fun with it if you wish to learn Go like me :)