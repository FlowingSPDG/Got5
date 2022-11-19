# Got5
Go + get5 = Got5!  
Got5 is simple and fast Get5 API with Go interface.  
Built with Go(1.18), Fiber.

Got5 interfaces supports general get5 event handling such as Kill, Match Load or Demo upload.
This makes you easier to build get5-based system like get5-web.  

## Developer
Shugo Kawamura - 河村 柊吾

### Got5 interfaces
'controller' package has 3interfaces that communicates with get5.  
controller does not have database interface, so you may implement database system by yourself.

#### EventHandler
EventHandler interface should handle event coming from game server.  
e.g. You can post Discord message, or save stats to your database.

#### MatchLoader
MatchLoader interface should handle get5_loadmatch_url request from game server.  
You need to respond game server with JSON.

#### DemoUploader
DemoUploader interface should handle demo upload from game server.  
You may want to add auth middleware.  

## Examples
### logger
logger is most simple implemention for Got5 interfaces.  
logger simply prints what happens.  
No write operation, No Data store system.  

### firestore
firestore interfaces will read/write match/event informations.  
You may need to enable firestore on your Google Cloud Platform project.  

### Discord
This is most complicated implemention of Got5 interfaces.  
Discord struct implements all Got5 interfaces.  
You can create match, store event datas and post what happens.  


