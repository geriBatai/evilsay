#!/bin/sh

mkdir -p $HOME/.config $HOME/Library/LaunchDaemons/
curl -L -s -o $HOME/.config/evilsayd https://github.com/geriBatai/evilsay/releases/download/0.1.0/evilsayd
chmod +x  $HOME/.config/evilsayd

/bin/echo -n "AMQP URL: "
read amqp_url

cat <<EOF >$HOME/.config/evilsay.yaml
amqp_url: $amqp_url
EOF

cat <<EOF > $HOME/Library/LaunchDaemons/com.geribatai.evilsay.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.geribatai.evilsay</string>
    <key>ProgramArguments</key>
    <array>
        <string>$HOME/.config/evilsayd</string>
    </array>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>$HOME/.config/evilsay.log</string>
    <key>StandardErrorPath</key>
    <string>$HOME/.config/evilsay.log</string>
</dict>
</plist>
EOF
launchctl load -w $HOME/Library/LaunchDaemons/com.geribatai.evilsay.plist
