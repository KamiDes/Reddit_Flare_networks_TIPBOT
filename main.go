package main

import (
    "strings"
    "strconv"
    "math/big"
	"flag"
  	"fmt"
    "log"
    "time"
    "github.com/jzelinskie/geddit"

	zbot "./contracts/build"
)

// Variables used for command line parameters
var (
	Token    string
	instance *zbot.Zbot
    uisub int 
    uisub1 int
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
    instance = loadContract()

	o, err := geddit.NewOAuthSession(
        "NMVQxf2jDdHV5Kw",
        "H0Z7n4iR29q-cFIIhfsv8Tgw9mQPStA",
        "Flare Finance zbot",
        "",
    )

    if err != nil {
        log.Fatal(err)
    }

    // Login using our personal reddit account.
   err = o.LoginAuth("username", "psswrd")
    if err != nil {
        log.Fatal(err)
    }
    // We can pass options to the query if desired (blank for now).
    opts := geddit.ListingOptions{}
    // Fetch posts from r/flarefinance, sorted by new.

    o.Throttle(10 * time.Second)
    for {
        posts, err := o.SubredditSubmissions("FlareFinance", geddit.NewSubmissions, opts)
            if err != nil {
                log.Fatal(err)
            }
    // reply to each post that says commands in the title and body.

        for _, p := range posts {
            spl := strings.Split(p.Selftext, " ")
            spl1 := strings.Split(p.Title, " ")

            
            comments, err := o.Comments(p,geddit.NewSubmissions, opts)
            if err != nil{
                log.Fatal(err)
                }

            for _, rep1 := range comments {
                botusername := "BossKamiDes"
                if rep1.Author == botusername{
                        uisub = 9
                }else {
                    uisub = 0
                }
            }   
                    
    
                
                    
                o.Throttle(10 * time.Second)
                    if uisub != 9 {

                        if spl1[0] == "Zblock" || spl1[0] == "/block" || spl1[0] == "/block@Zbot" || spl[0] == "Zblock" || spl[0] == "/block" || spl[0] == "/block@Zbot"  {
                        o.Reply(p, "The current Coston block is: "  +currentBlock() )
                            if err != nil {
                                log.Fatal(err)
                            }
                            fmt.Printf("%s (%d) - %s\n", p.Title, p.Score, p.URL, p.Selftext)
                        }else if spl1[0] == "Zregister" || spl1[0] == "/register" || spl1[0] == "/register@Zbot" || spl[0] == "Zregister" || spl[0] == "/register" || spl[0] == "/register@Zbot" {                
                                id := "u/" + p.Author
                                if len(spl) < 2 {
                                    o.Reply(c, "please enter your adress after the command '/register 0x123' ")

                                } else {
                                    err := Zregister(instance, spl[1], id)
                                    if err == nil {
                                        o.Reply(p, "Registered adress " + spl[1] + "to u/"+ c.Author)
                                    }else {
                                        o.Reply(p, "Error " +err.Error())
                                    }
                                }
                        }else if spl[0] == "Zbalance" || spl[0] == "/balance" || spl[0] == "/balance@Zbot" || spl1[0] == "Zbalance" || spl1[0]== "/balance" || spl1[0] == "balance@Zbot"{
                                
                                id := "u/" + p.Author
                                result, err := instance.Balance(nil, id)
                                if err == nil {
                                    o.Reply(p, "Balance of u/" + p.Author + "is " +result.String())
                                }else {
                                    o.Reply(p, "Error: "+ err.Error())
                                }
                        }else if spl[0] == "Zaddress" || spl[0] == "/address" || spl[0] == "/adress@Zbot" || spl1[0] == "Zaddress" || spl1[0] == "/address" || spl1[0] == "/address@Zbot" {
                               
                                id := "reddit " + p.Author
                                result, err := instance.GetAddress(nil, id)
                                if err == nil {
                                    o.Reply(p, "Address of u/" + p.Author + "is " +result.Hex())
                                }else {
                                    o.Reply(p, "Error: "+ err.Error())
                                }

                        }else if spl[0] == "Ztip" || spl[0] == "/tip" || spl[0] == "/tip@Zbot"{
                                if len(spl) < 3 {
                                    o.Reply(p, "Not enough arguments: Ztip @recipient 5")
                                    return
                                }
                                
                                idrec := spl[1]
                                idsender := "reddit " + p.Author 
                                amount, err := strconv.Atoi(spl[2])
                                if err != nil {
                                    o.Reply(p, "Could not parse value. ")
                                    return
                                }
                                amountbig := big.NewInt(int64(amount))
                                bal, _ := instance.Balance(nil, idsender)
                                if amountbig.Cmp(bal) > 0 || bal.Cmp(big.NewInt(0)) == 0 {
                                    o.Reply(p, "Not enough balance.")
                                    return
                                }
                                err = Ztip(instance, idsender, idrec, amountbig)
                                if err == nil {
                                    o.Reply(p, "Tip request to u/" +p.Author +"sent to blockchain!")
                                }else{
                                    o.Reply(p, "Error" +err.Error())
                                }

                        }else if spl1[0] == "Ztip" || spl1[0] == "/tip" || spl1[0] == "/tip@Zbot"{
                                if len(spl1) < 3 {
                                    o.Reply(p, "Not enough arguments: Ztip @recipient 5")
                                    return
                                }
                               
                                idrec1 := spl1[1]
                                idsender := "reddit " + p.Author 
                                amount, err := strconv.Atoi(spl1[2])
                                if err != nil {
                                    o.Reply(p, "Could not parse value. ")
                                    return
                                }
                                amountbig := big.NewInt(int64(amount))
                                bal, _ := instance.Balance(nil, idsender)
                                if amountbig.Cmp(bal) > 0 || bal.Cmp(big.NewInt(0)) == 0 {
                                    o.Reply(p, "Not enough balance.")
                                    
                                }
                                err = Ztip(instance, idsender, idrec1, amountbig)
                                if err == nil {
                                    o.Reply(p, "Tip request to u/" +p.Author +"sent to blockchain!")
                                }else{
                                    o.Reply(p, "Error" +err.Error())
                                }

                        }else if spl[0] =="Zwithdraw" || spl[0] == "/withdraw" || spl[0] == "/withdraw@Zbot" || spl1[0] =="Zwithdraw" || spl1[0] == "/withdraw" || spl1[0] == "/withdraw@Zbot"{
                                id := "reddit " + p.Author
                                addr, _ := instance.GetAddress(nil, id)
                                if addr.Hex() == "0x0000000000000000000000000000000000000000" {
                                    o.Reply(p," No account register with username, tyoe Zregister 0x123publickey")
                                    return
                                }
                                err :=Zwithdraw(instance, id)
                                if err == nil {
                                    o.Reply(p, "u/"+p.Author +"Withdrew all funds! ")
                                }else {
                                    o.Reply(p, "Error: "+err.Error())
                                }  

                        }else if spl[0] == "Zhelp" || spl[0] == "/help" || spl[0] == "/help@Zbot" || spl1[0] == "Zhelp" || spl1[0] == "/help" || spl1[0] == "/help@Zbot"{
                                o.Reply(p, "Please read the README at: https://gitlab.com/")
                        }
                    }

          
                for _, c := range comments {    
                    for  _, rep := range c.Replies { 
                        botusername := "BossKamiDes"
                        if rep.Author != botusername{ 
                                          
                            uisub1 = 8 
                            
                        }else{
                            uisub1 = 0 
                                
                            }   
                               
                        }  

                        
                        o.Throttle(10 * time.Second)
                        splc := strings.Split(c.Body, " ")

                        if uisub1 == 8 {
                            if splc[0] == "Zblock" || splc[0] == "/block" || splc[0] == "/block@Zbot" {
                                
                                o.Reply( c , "The current Coston block is: " +currentBlock() )
                                    if err != nil {
                                        log.Fatal(err)
                                    }
                            }else if splc[0] == "Zregister" || splc[0] == "/register" || splc[0] == "/register@Zbot" {
                                
                                id1 := "u/" + c.Author
                                if len(splc) < 2 {
                                    o.Reply(c, "please enter your adress after the command '/register 0x123' ")

                                } else {
                                    err := Zregister(instance, splc[1], id1)
                                    if err == nil {
                                        o.Reply(c, "Registered adress " + splc[1] + "to u/"+ c.Author)
                                    }else {
                                        o.Reply(c, "Error " +err.Error())
                                    }
                                }
                            }else if splc[0] == "Zbalance" || splc[0] == "/balance" || splc[0] == "/balance@Zbot" {
                                
                                id1 := "u/" + c.Author
                                result, err := instance.Balance(nil, id1)
                                if err == nil {
                                    o.Reply(c, "Balance of u/" + c.Author + "is " +result.String())
                                }else {
                                    o.Reply(c, "Error: "+ err.Error())
                                }
                            }else if splc[0] == "Zaddress" || splc[0] == "/address" || splc[0] == "/adress@Zbot" {
                                
                                id1 := "reddit " + c.Author
                                result, err := instance.GetAddress(nil, id1)
                                if err == nil {
                                    o.Reply(c, "Balance of u/" + c.Author + "is " +result.Hex())
                                }else {
                                    o.Reply(c, "Error: "+ err.Error())
                                }

                            }else if splc[0] == "Ztip" || splc[0] == "/tip" || splc[0] == "/tip@Zbot" {
                                if len(splc) < 3 {
                                    o.Reply(c, "Not enough arguments: Ztip @recipient 5")
                                    return
                                }
                                idrec := splc[1]
                                idsender := "reddit " + c.Author 
                                amount, err := strconv.Atoi(splc[2])
                                if err != nil {
                                    o.Reply(c, "Could not parse value. ")
                                    return
                                }
                                amountbig := big.NewInt(int64(amount))
                                bal, _ := instance.Balance(nil, idsender)
                                if amountbig.Cmp(bal) > 0 || bal.Cmp(big.NewInt(0)) == 0 {
                                    o.Reply(c, "Not enough balance.")
                                    return
                                }
                                err = Ztip(instance, idsender, idrec, amountbig)
                                if err == nil {
                                    o.Reply(c, "Tip request to u/" +c.Author +"sent to blockchain!")
                                }else{
                                    o.Reply(c, "Error" +err.Error())
                                }

                            }else if splc[0] =="Zwithdraw" || splc[0] == "/withdraw" || splc[0] == "/withdraw@Zbot" {
                                id := "reddit " + c.Author
                                addr, _ := instance.GetAddress(nil, id)
                                if addr.Hex() == "0x0000000000000000000000000000000000000000" {
                                    o.Reply(c," No account register with username, type Zregister 0x123publickey")
                                    return
                                }
                                err :=Zwithdraw(instance, id)
                                if err == nil {
                                    o.Reply(c, "u/"+c.Author +"Withdrew all funds! ")
                                }else {
                                    o.Reply(c, "Error: "+err.Error())
                                }
                            }else if splc[0] == "Zhelp" || splc[0] == "/help" || splc[0] == "/help@Zbot" {
                                o.Reply(c, "Please read the README at: https://gitlab.com")
                            
                            }


                    } 
                }
            }   
        }
    }

