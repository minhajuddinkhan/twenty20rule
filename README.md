# twenty20rule

Sends a desktop notification every twenty minutes based on the 20 20 rule.
## Setup

![alt text](https://image.ibb.co/cRX05b/Screenshot_from_2017_10_12_21_23_59.jpg)

make sure  your ```$PATH``` has ```$GOPATH/bin```

``` $ go get github.com/minhajuddinkhan/twenty20rule ```

``` $ twenty20rule & ```

### Initialize everytime PC reboots?

Add ``` twenty20rule &  ``` in ```/etc/rc.local``` before ```exit 0``` statement. (Linux)

(open for how do to this in mac? lol)


## Golang not configured?

 Download ![Precompiled binaries](https://github.com/minhajuddinkhan/twenty20rule/releases/tag/0.1.0) 

## Description 
It's no secret that sitting all day damages your body, but figuring out a system to counteract that for yourself is tough. The New York Times suggests one simple rule you can employ is a variation on the 20-20 rule to reduce eyestrain, but with movement instead.

![alt text](http://www.anthro.com/getmedia/b54b97a8-b2b4-401c-b143-ef15922b003b/20-20-20-ergo-tip?width=600&height=368&ext=.jpg)
