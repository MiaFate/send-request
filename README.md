### 1. build the binary
```bash
 go build -o sendreq main.go
```

### 2. add your token and url into sample.env file and rename it to .env
```
BEARER=your_token_here
URL=your_url_here
```
### 3. finished.txt must have just an empty line

### 4. load your products into products.txt
example:
```
product1
product2
product3
```

### 5. add path to this folder in scriptrunner.sh line 8
```bash
cd /Users/<your_user>/<path>/send-request
```

### 5. give permission to setup.sh
```bash
chmod +x setup.sh
```

### 6. run setup.sh
```bash
./setup.sh
```


if you want to run the script manually, you can run the scriptrunner.sh file
```bash
### cronjob file
create a new crontab with `crontab -e`
and add the following line to run the script every minute
```bash
*/1 * * * * /Users/miafate/programacion/golang/src/github.com/miafate/goscripts/send-request/scriptrunner.sh
``` 
save and exit with `:wq`