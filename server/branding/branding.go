package branding

import (
	"fmt"
	"time"
)

var Name = "Aegis Terminal"
var Copyright = fmt.Sprintf("Copyright © 2020-%d dushixiang, All Rights Reserved.", time.Now().Year())
var Banner = `    ___       ___   
   /\__\     /\  \  
  /:| _|_    \:\  \ 
 /::|/\__\   /::\__\
 \/|::/  /  /:/\/__/
   |:/  /   \/__/   
   \/__/            `
var Version = `v1.3.9`
var Hi = Banner + Version
