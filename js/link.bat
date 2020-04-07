#mklink /d node_modules A:\node_modules
rmdir A:\radar_cache
mkdir A:\radar_cache
rmdir .cache
mklink /d .cache A:\radar_cache
rmdir A:\radar_arts
mkdir A:\radar_arts
rmdir artifacts
mklink /d artifacts A:\radar_arts
