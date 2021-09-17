# gorestapi
Quick start with  GO lang to implement rest api and various operations. 

## Description
Rest api exposes GET AND POST methods to user. 

Item structure is as below -
```golang
 type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
```


### GET all items
![get](https://github.com/jai2036/gorestapi/blob/master/IMAGES/GO-rest-get-api.png)

### POST add items to cache
![post](https://github.com/jai2036/gorestapi/blob/master/IMAGES/GO-rest-post-api.png)


## Dependecies
1. Docker 

## Executing
1. Create a project direcotry (optional) 
```bash
$ mkdir -p Projects 
$ cd Projects
```
2. Clone the repo
```bash
~/Projects/$ gh repo clone jai2036/gorestapi
```
3. Go to app directory (~/Projects/gorestapi/) and build the dockerfile. 
```bash
~gorestapi/$ docker build -t jai/myapp .
```
4. Run the image.
```bash
# list images and copy the id
~gorestapi/$ docker images
# Run your image 
~gorestapi/$ docker run -d -p 8080:8080 <image id>
```
