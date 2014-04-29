package steamail_test

import (
	"bytes"
	"net/mail"
	"testing"

	. "github.com/garslo/steambot/steamail"
)

func TestWeCanExtractTheCode(t *testing.T) {
	extractor := NewSteamCodeExtractor()
	msg, err := mail.ReadMessage(bytes.NewBufferString(ExampleEmail))
	if err != nil {
		t.Errorf("Could not read message: %v", err)
	}
	code, err := extractor.ExtractCode(msg)
	if err != nil {
		t.Errorf("Could not extract code: %v", err)
	}
	if code != "4QJY8" {
		t.Errorf("Bad code; got '%s', wanted '4QJY8'", code)
	}
}

const ExampleEmail = `Delivered-To: example@gmail.com
Received: by 10.76.173.7 with SMTP id bg7csp103220oac;
        Sun, 27 Apr 2014 18:18:31 -0700 (PDT)
X-Received: by 10.66.189.106 with SMTP id gh10mr22522081pac.31.1398647911508;
        Sun, 27 Apr 2014 18:18:31 -0700 (PDT)
Return-Path: <noreply@steampowered.com>
Received: from smtp02.steampowered.com (smtp02.steampowered.com. [208.64.202.38])
        by mx.google.com with ESMTPS id nl9si1446004pbc.8.2014.04.27.18.18.31
        for <example@gmail.com>
        (version=TLSv1 cipher=RC4-SHA bits=128/128);
        Sun, 27 Apr 2014 18:18:31 -0700 (PDT)
Received-SPF: pass (google.com: domain of noreply@steampowered.com designates 208.64.202.38 as permitted sender) client-ip=208.64.202.38;
Authentication-Results: mx.google.com;
       spf=pass (google.com: domain of noreply@steampowered.com designates 208.64.202.38 as permitted sender) smtp.mail=noreply@steampowered.com;
       dkim=pass header.i=@steampowered.com
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/relaxed; d=steampowered.com; s=smtp;
	h=Date:Message-Id:Content-Type:Subject:MIME-Version:Reply-To:From:To; bh=tg4SknCC1y5zQlhBam200ED9eGk3OZmGm9faDxJp8Zg=;
	b=eLDEA/Zd5CD3lzhD25H2qgi/tOLTtKKa5iOnbZNgdBW80SvvLu4vMAODFi7Gie3eQHC8cZRzPCoElyWMnSkl/TqR++Tb2KI3aCpDuj7oYmIDqC3+BQzyOv8czRwBHSAtKI1V6u4HGrnT4/D/SwitXOo+0tQwuRIUnnNazLJPjcA=;
Received: from [10.3.3.4] (helo=valvesoftware.com)
	by smtp02.steampowered.com with smtp (Exim 4.76)
	(envelope-from <noreply@steampowered.com>)
	id 1WeaC4-0002vc-VW
	for example@gmail.com; Sun, 27 Apr 2014 18:17:16 -0700
To: example@gmail.com
From: Steam Support <noreply@steampowered.com>
Reply-To: <noreply@steampowered.com>
X-Steam-Message-Type: Account Information Confirmation
MIME-Version: 1.0
Subject: Your Steam account: Access from new computer
Content-Type: multipart/alternative;
 boundary="------------060908020109090601040503"
Message-Id: <E1WeaC4-0002vc-VW@smtp02.steampowered.com>
Date: Sun, 27 Apr 2014 18:17:16 -0700

This is a multi-part message in MIME format.
--------------060908020109090601040503
Content-Type: text/plain; charset=UTF-8; format=flowed
Content-Transfer-Encoding: 7bit


Dear steam_username,

We've received a request to access your Steam account from the Steam Client on
a new computer located at IP address: 192.168.1.1
Our records show this IP address is in ANYTOWN, US

The following code is only for logging into the Steam Client.
Do not enter this code into a web browser. If this email was generated as a result
of you entering your account name and password into a web site, that site may be
malicious, and we recommend that you change your password immediately.
Click here for more information: https://steamcommunity.com/actions/ReportSuspiciousLogin?stoken=0c69f19eb352916272e80d38c8eea9ee755bab1241f

To complete the login, enter the following special access code into the
authorization dialog in the Steam Client before trying to log in again: 4QJY8

If you did not attempt this action, please change your password immediately.

Thanks for helping us maintain the security of your account.


The Steam Support Team
https://support.steampowered.com



==============
This notification has been sent to the email address associated with your Steam account.
For information on Valve's privacy policy, visit http://www.valvesoftware.com/privacy.htm.
This email message was auto-generated. Please do not respond.

Â© Valve Corporation. All rights reserved. All trademarks are property of their respective
owners in the US and other countries.


--------------060908020109090601040503
Content-Type: text/html; charset=UTF-8
Content-Transfer-Encoding: 7bit

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<html>
	<head>
	<meta content="text/html;charset=UTF-8" http-equiv="Content-Type">
	</head>
		<body bgcolor="#ffffff" text="#000000" LINK="#517ba4" ALINK="#517ba4" VLINK="#517ba4">

		<span style="font-size: 10px; background: #000000; color: #999999; font-family: Trebuchet MS,Verdana,Arial,Helvetica,sans-serif; lang="x-western">
		<img src="http://storefront.steampowered.com/v/img/email/STEAM_mail.jpg" alt="&nbsp;STEAM&nbsp;" width="430" height="48" border="0"><BR>
		</span>

			<div style="width: 430px; font-size: 11px; color: #333333; font-family: Trebuchet MS,Verdana,Arial,Helvetica,sans-serif;  lang="x-western">
			<BR>
			Dear steam_username, <br>
			<br>
			We've received a request to access your Steam account from the Steam Client on <br>
			a new computer located at IP address: <b style="font-size: 15px;">192.168.1.1</b> <br>
			Our records show this IP address is in <b style="font-size: 15px;">ANYTOWN, US</b><br>

			<p>The following code is only for logging into the Steam Client.<br>
			<b>Do not enter this code into a web browser.</b> If this email was generated as a result <br>
			of you entering your account name and password into a web site, that site may be <br>
			malicious, and we recommend that you change your password immediately. <br>
			Click here for more information:</p>
			<h2><a href="https://steamcommunity.com/actions/ReportSuspiciousLogin?stoken=d38c8eea9ee755bab1241f">Verify Login Location</a></h2>

			<p>To complete the login, enter the following special access code into the <br>
			authorization dialog in the Steam Client before trying to log in again: </p>
			<h3>4QJY8</h3>
			<p>If you did not attempt this action, please change your password immediately. <br>
			<p>Thanks for helping us  maintain the  security of your account.<BR>
			<p>The Steam Support Team<br>
			<a href="https://support.steampowered.com">https://support.steampowered.com</a> <br>
			<p></p>
			<br>
			<br>
				<span style="color: #333333; font-size: 9px; font-family: Trebuchet MS, Verdana, Arial, Helvetica, sans-serif;">This notification has been sent to the email address associated 	with your Steam account. <br>
				For information on Valve's privacy policy, visit <a href="http://www.valvesoftware.com/privacy.htm">http://www.valvesoftware.com/privacy.htm</a>. <br>
				This email message was auto-generated. Please do not respond.</span><br>
			<hr color="#666666" align="left" width="420" size="1" noshade>
				<table width="420" border="0" cellspacing="0" cellpadding="0">
					<tr>
						 <td><span style="font-size: 10px; background: #FFFFFF; color: #333333; font-family: Trebuchet MS,Verdana,Arial,Helvetica,sans-serif; lang="x-western"><img src="http://storefront.steampowered.com/v/img/email/VALVE_mail.gif" alt="VALVE" width="91" height="25"></span></td>
						 <td width="320"><span style="color: #333333; font-size: 9px; font-family: Trebuchet MS, Verdana, Arial, Helvetica, sans-serif;">&copy; Valve Corporation. All rights reserved. All trademarks are property of their respective owners in the US and other countries.</span></td>
					</tr>
				</table>
			</div>
		</body>
</html>


--------------060908020109090601040503--`
