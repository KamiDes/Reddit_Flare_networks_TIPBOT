# Zbot

Reddit Bot for the Flare network.

Commands work on the title, body and comments of posts, Feel free to include a message after the commands but not before

If you have received tips it is easy to start tipping, just type
```
/tip @recipient 100000000000000000
```
to tip 0.1 FLR.

Use /balance to check your balance.

To withdraw you first have to register an address:
```
/register 0x123ADDRESS
```
Then simply
```
/withdraw
```

The contract address is: 0xCBB39D17B5a45198d57586e86EF05C715a3ce025

To deposit either send money to the contract address with your registered address or run the deposit script in the deposit folder, first change the ID to your discord ID: https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-

Then run it like so:
```
SECRET=123YOURPRIVATEKEY go run deposit.go
```

After updating Zbot.sol in contracts folder run:
```
rm -r build
solc --abi --bin Zbot.sol -o build
abigen --bin=./build/Zbot.bin --abi=./build/Zbot.abi --pkg=zbot --out=build/Zbot.go
SECRET=PRIVATEKEY go run deploy.go
```
Then copy the new contract address to flare.go and deposit.go

To run the bot do
```
SECRET=123PRIVATEKEYOFCONTRACTOWNER TELEGRAM=TELEGRAMBOTTOKEN REDSECRET=THE-SECRET REDPASSWORD=YourPassword go run *.go -t Abc.discordbot.token
```
