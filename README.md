<p align="center">
  <img src="https://64.media.tumblr.com/bdba59dac0ceb4ccfbee778be8f96b69/tumblr_ow6999EESx1qfu8poo1_400.gif" width="250" height="250">
</p>

    >> two scripts written in golang that will help you recognize dependency confusion.    
---
```
> to check if the sites have the package.json or requirements.txt file:

  >> ./check -w sites.txt
  
> to check if libs exist externally:

  >> ./confusion -w libs.txt -t TYPE (gem, pip or npm)
```
</p>
<p align="center">
  <img src="https://user-images.githubusercontent.com/44043159/135767961-fc43ca2f-04f1-43ec-a1eb-d8c58248ac44.png" width="500" height"500">
  <img src="https://user-images.githubusercontent.com/44043159/135768086-3dd610b4-6038-4728-94e9-2ba86c9eda2b.png" width="500" height"500">
</p>

```bash
git clone https://github.com/march0s1as/dep-confusion
cd dep-confusion
go get -v github.com/fatih/color
go build check.go
go build confusion.go
enjoy =]
```
