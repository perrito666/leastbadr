package main

func main(){
	return os.Exit(runner())
}

type config struct{
	templatePath string
	picturePath string
	buildHTML bool
	buildThumbs bool
	syncUp bool
}

func runner() int {
	c := &config{}
	&c.templatePath = flag.String("template", "", "Path to the folder containing the html templates")
	&c.picturePath = flag.String("photos", "", "Path to the folder containing the photos")
	&c.buildHTML = flag.Bool("html", false, "Build local HTML")
	&c.buildThumbs = flag.Bool("thumbs", false, "Build local Thumbnails")
	&c.syncUp = flag.Bool("sync", false, "Sync with remote storage")
	return 0
}