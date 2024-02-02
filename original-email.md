# Original communication of file format

``` SMTP
This package provides a Go implementation of the VSA file format.

Return-path: <sales@lifeape.com>
Original-recipient: rfc822;tmornini@me.com
Received: from ci74p00im-qukt09072702.me.com by p116-mailgateway-smtp-746c496c44-6mknn (mailgateway 2403B92)
	with SMTP id 4f24e911-87c4-4315-8f0b-f8245274ec19 
	for <tmornini@me.com>; Tue, 2 Jan 2024 18:52:56 GMT
X-Apple-MoveToFolder: INBOX 
X-Apple-Action: MOVE_TO_FOLDER/INBOX
X-Apple-UUID: 4f24e911-87c4-4315-8f0b-f8245274ec19
Received: from sender4-of-o50.zoho.com (sender4-of-o50.zoho.com [136.143.188.50])
	by ci74p00im-qukt09072702.me.com (Postfix) with ESMTPS id 11C766C0012C
	for <tmornini@me.com>; Tue,  2 Jan 2024 18:52:53 +0000 (UTC)
X-ICL-SCORE: 3.332033030043
X-ICL-INFO: GAtbVUseBVFHSVVESAMGUkFKRFcUWUIPAApbVRYSFhEAREQZF15TQFUcAkpaQ1cOEAomGxFWUwMF
 HEgUF10UQhMdW1UUWVAHFAkDWRtfW0BVCg8HRRIHUUNXV0NLHgdaTVdTR1oQXgcZFltVC1VYBhAL
 UloXVhsNQElIDRdYWUwWFgtVWEBCEEhbFRIWVFMQQVQJEVVfA1JbAwkfFx5VDRhbRhMcDRQOHB8D
 FghVGAEaFERXFVlSX1dFV08bU1RaQE8AQk4eUlZASQMVTAQFVUxMUUVWBQJWQE1VRE0CDlNCQVVB
 TggDUDUVDxEdUUYHWxoJGkYSFhAWREQDFV9EDBwXDzcVVRgBGhRE
Authentication-Results: bimi.icloud.com; bimi=skipped reason="insufficient dmarc"
X-ARC-Info: policy=fail; arc=pass; id=mx.zohomail.com
Authentication-Results: arc.icloud.com; arc=pass
Authentication-Results: dmarc.icloud.com; dmarc=none header.from=lifeape.com
X-DMARC-Info: pass=none; dmarc-policy=(nopolicy); s=u0; d=u0; pdomain=lifeape.com
X-DMARC-Policy: none
Authentication-Results: dkim-verifier.icloud.com;
	dkim=pass (1024-bit key) header.d=lifeape.com header.i=sales@lifeape.com header.b=I77INjus
Authentication-Results: spf.icloud.com; spf=pass (spf.icloud.com: domain of sales@lifeape.com designates 136.143.188.50 as permitted sender) smtp.mailfrom=sales@lifeape.com
Received-SPF: pass (spf.icloud.com: domain of sales@lifeape.com designates 136.143.188.50 as permitted sender) receiver=spf.icloud.com; client-ip=136.143.188.50; helo=sender4-of-o50.zoho.com; envelope-from=sales@lifeape.com
ARC-Seal: i=1; a=rsa-sha256; t=1704221570; cv=none; 
	d=zohomail.com; s=zohoarc; 
	b=REJZtaTHEqDmLQlZk0xe3fMVAXNR1TPwAL5ohpAaprUic6a9+PV1gaC0lzN6vYaBPKF48eNgQTHQyOjp1EDibRhMqh/I0x2QADKOioDKYFS+zTjLF6uIM74HRNSJDYxNxfLMg7odSG1z6ke7niFg5rQX2YsNMbeUjVh87xF4CyQ=
ARC-Message-Signature: i=1; a=rsa-sha256; c=relaxed/relaxed; d=zohomail.com; s=zohoarc; 
	t=1704221570; h=Content-Type:Date:Date:From:From:In-Reply-To:MIME-Version:Message-ID:References:Subject:Subject:To:To:Message-Id:Reply-To:Cc; 
	bh=uJYs/gO31qECRoMF0A6TEank793esQK8D+9Qoj4UCnE=; 
	b=h/019/JftuRXeSQhvwRp7SuCCDYgHGM7NGQAG5wN/kjCeLqll46b5/SU5eVwgRoivZJThFNKj1UH6IMH3myPPPL6N5T4mxA8gizTLkvu/yYrFFpJ3+v/PmDw8cbTpcwHIA9rRAG7I9zHB3e5vuVZXO7DEv90dJx4jPKjLvujWPs=
ARC-Authentication-Results: i=1; mx.zohomail.com;
	dkim=pass  header.i=lifeape.com;
	spf=pass  smtp.mailfrom=sales@lifeape.com;
	dmarc=pass header.from=<sales@lifeape.com>
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/relaxed; t=1704221570;
	s=default; d=lifeape.com; i=sales@lifeape.com;
	h=Date:Date:From:From:To:To:Message-Id:Message-Id:In-Reply-To:References:Subject:Subject:MIME-Version:Content-Type:Reply-To:Cc;
	bh=uJYs/gO31qECRoMF0A6TEank793esQK8D+9Qoj4UCnE=;
	b=I77INjusJsG4KvkWBIFLzlR864JEVRI2lF2QKUFRadyMAg2wk02ZgcC+xANOW01d
	jKJJUUd203Ho4RLM2t0Xzj2z+qZedJ0hwHJo4oboyaqCUiqRNDzfGuvkslaiGGE7ZjK
	N6hoESHRgIEay2pS9gxA/QPe3UuEZF0aq+1yT8O0=
Received: from mail.zoho.com by mx.zohomail.com
	with SMTP id 1704221568886815.6670060217856; Tue, 2 Jan 2024 10:52:48 -0800 (PST)
Date: Tue, 02 Jan 2024 13:52:48 -0500
From: Sales <sales@lifeape.com>
To: "Tom Mornini" <tmornini@me.com>
Message-Id: <18ccb856f56.d450eb44379572.5445433528178366852@lifeape.com>
In-Reply-To: <03261038-87EF-49D0-AF73-94DECD700D77@me.com>
References: <E2DB12C6-9772-4156-B89A-AB83F588AAE9@lifeape.com> <03261038-87EF-49D0-AF73-94DECD700D77@me.com>
Subject: Re: LifeApe "Did you reverse engineer the VSA format?"
MIME-Version: 1.0
Content-Type: multipart/alternative; 
	boundary="----=_Part_1160233_750255476.1704221568855"
Importance: Medium
User-Agent: Zoho Mail
X-Mailer: Zoho Mail
X-CLX-Shades: Deliver
X-MANTSH: 1TFkXGxoaHxEKWUQXbRlNEk1uTHoYXh8RCllNF25PRkNcT1gRCl9ZFxsaGhEKX00
 XZEVETxEKWUkXHR9xGwYbHxp3BhsaEgYaBhoGHxIGBxsdG3EaEB8adwYaBhoGGRoGGgYaBhpxG
 hAadwYaEQpZXhdsbHkRCkNOFxlSckxYTEB1XG4dWmB7UH1mfFx5bntmGU51YV5MX2h/EQpYXBc
 ZBBoEHxoFGxoaBBIYBBsfEgQYHBAbHhofGhEKXlkXSVNNa0IRCk1cFxsbHREKTFoXaGhrQWsRC
 k1OF2hrEQpMRhdiTWsRCkNaFxsZHAQbHhkEGxISBB8aEQpCXhcbEQpCXBcbEQpeThcbEQpCSxd
 jEh1dbh5TTR9kSREKQkkXYxIdXW4eU00fZEkRCkJFF2ZCeElDBRx6Wk1gEQpCThdjEh1dbh5TT
 R9kSREKQkwXYH8BQ3sFQHobbFgRCkJsF2Ble0gYcFJcX21oEQpCQBdrb3xgGxMeRUt6eREKQlg
 XYllAZBMcbUt8en4RCkJ4F2ddU19PQ0waWE1SEQpNXhcbEQpaWBcYEQpwaBduRUxfAUBOaUBsR
 RAZGhEKcGgXbURPYWJ7WWVORXgQGRoRCnBoF3pwfUR8fVliZnt7EBkaEQpwaBdkZm9NUmNcQRM
 eTBAZGhEKcGgXYn8BaVhoX1AFQEQQGRoRCnBoF29/e3NiaEh+QWVHEBkaEQpwaBduG00Zf31ua
 xtOchAZGhEKcGgXaG5pHU0abmduGF0QGRoRCnB/F2haUFhbbV8bTHhyEBkaEQpwfRdoWlBYW21
 fG0x4chAZGhEKcGwXZEZOWmhORUxhTmwQBxoEHBkRCnBMF299Q35ZUGFsQ2BaEBkaEQptfhcbE
 QpYTRdLEQ==
X-Proofpoint-ORIG-GUID: 3xXfrfj_vD7pJQzWLVvSDQL3d_KtfuBU
X-Proofpoint-GUID: 3xXfrfj_vD7pJQzWLVvSDQL3d_KtfuBU

------=_Part_1160233_750255476.1704221568855
Content-Type: text/plain; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

Hi Tom,



Here you go. Hopefully this makes sense, I wrote it about 8 years ago.



-------------------------------=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0

Note: (#) =3D number of bytes

=C2=A0

DATA

// Version - DATA(12)

// Level - SIZE(1) DATA(S)

// Options - SIZE(1) DATA(S)

// Email - SIZE(1) DATA(S)

// # of events - DATA(4)

// Other - DATA(4)

=C2=A0

EVENTS

// CEventBarLinear or
CEventBarPulse - SIZE(1) 00 DATA(S)

// Events - TRAK(2) STRT(2)
0000 ENDT(2) 0000 SPOS(2) 0000 EPOS(2) 0000 SIZE(1) TEXT(S) DATA(12 or 16) =
XX
XX

// DATA(12)
=3D Linear

//
DATA(16) =3D Pulse

=C2=A0=C2=A0=C2=A0 =C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=
=C2=A0=C2=A0 //
XX XX decipher

=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=A0=C2=
=A0=C2=A0=C2=A0 // FF FF =3D New Event Type

//followed by CEventBarLinear or CEventBarPulse -
SIZE(1) 00 DATA(S)

//
01 80 =3D Same as first event type

//
30 87 =3D Other event type

//
00 00 =3D Last event

=C2=A0

AUDIO FILES

//# of Files - DATA(2) 0000 (IF
FILES)

// Audio File - SIZE(1)
FILE(S) SIZE(1) AUDIO_DEVICE(S)

=C2=A0

VIDEO FILES

// # of Files - DATA(2)
0000(IF FILES)

// Video File - SIZE(1)
FILE(S) SIZE(1) AUDIO(S) SIZE(1) MONITOR(S) FULLSCREEN(1) XOFFSET(2) 0000
YOFFSET(2)

=C2=A0

// Other - DATA(14)

=C2=A0

TRACK SETTINGS

// # of Tracks - DATA(2) 0000

// Tracks - SIZE(1) TEXT(S)
ADDR(2) 0000 CNTR(1) DATA(11) +VAL(2) 0000 -VAL(2) 0000 NEUT(2) 0000 ENBL(1=
) 00
0000 SIZE(1) PORT(S) DATA(12)



-----------------------------------



Thanks,

Nelson









---- On Wed, 27 Dec 2023 17:48:37 -0500 Tom Mornini <tmornini@me.com> wrote=
 ---



Hey, thanks for getting back to me.



Yes, I=E2=80=99d love to know what you know about the format!



-- Tom Mornini

-- Sent from iPhone

-- Forgive brevity and typos





On Dec 27, 2023, at 2:03=E2=80=AFPM, Sales <mailto:sales@lifeape.com> wrote=
:



=EF=BB=BFHi Tom,



Thought I had replied to this but maybe not. I know Jerry very well, he and=
 I work together on a number of projects. I did reverse engineer the VSA fo=
rmat and am willing to share it if you still need it.=C2=A0



Thanks,

Nelson




On Nov 3, 2023, at 4:41=E2=80=AFPM, Tom Mornini <mailto:tmornini@me.com> wr=
ote:



=EF=BB=BFSure.



My buddy runs SkullTronix and I=E2=80=99m helping him build a small control=
ler board to operate his animatronic products without the need to be tether=
ed to a computer.



I would love to open source the documentation, and provide a Go language im=
plementation, but that=E2=80=99s subject to your approval, of course.



I very much appreciate your consideration, this would save me so much time =
and hassle! =F0=9F=99=8F



-- Tom Mornini

-- Sent from iPhone

-- Forgive brevity and typos





On Nov 3, 2023, at 4:29 AM, Sales <mailto:sales@lifeape.com> wrote:



=EF=BB=BF

Hi Tom,



I was able to determine the VSA file format and would be willing to share i=
t. Can I ask for a bit more information about the project?



Thanks,

Nelson












---- On Tue, 31 Oct 2023 16:33:56 -0400 Sales @ LifeApe <mailto:sales@lifea=
pe.com> wrote ---



From: Tom Mornini
 Email: mailto:tmornini@me.com
 Subject: Did you reverse engineer the VSA format?
Message Body:
 Hello there.
I'm working on a software project and the person I'm building it for would =
love it to be able to read VSA file formats directly.

Do you have documentation of that format that you'd be willing to share?

If not, would you be willing to sell it?

I appreciate your time, I'd really rather not have to figure this out mysel=
f.=20

--
 This e-mail was sent from a contact form on LifeApe (https://lifeape.com)
------=_Part_1160233_750255476.1704221568855
Content-Type: text/html; charset="UTF-8"
Content-Transfer-Encoding: quoted-printable

<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"><html><head>=
<meta content=3D"text/html;charset=3DUTF-8" http-equiv=3D"Content-Type"></h=
ead><body ><div style=3D"font-family: Verdana, Arial, Helvetica, sans-serif=
; font-size: 10pt;"><div>Hi Tom,<br></div><div><br></div><div>Here you go. =
Hopefully this makes sense, I wrote it about 8 years ago.<br></div><div><br=
></div><div>-------------------------------&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b=
r></div><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span =
style=3D"font-size:10.0pt">Note: (#) =3D number of bytes</span><br></p><p c=
lass=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"fon=
t-size:10.0pt">&nbsp;</span><br></p><p class=3D"" style=3D"margin-top: 0px;=
 margin-bottom: 0px;"><u><span style=3D"font-size:10.0pt">DATA</span></u><b=
r></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span st=
yle=3D"font-size:10.0pt">// Version - DATA(12)</span><br></p><p class=3D"" =
style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.=
0pt">// Level - SIZE(1) DATA(S)</span><br></p><p class=3D"" style=3D"margin=
-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">// Options=
 - SIZE(1) DATA(S)</span><br></p><p class=3D"" style=3D"margin-top: 0px; ma=
rgin-bottom: 0px;"><span style=3D"font-size:10.0pt">// Email - SIZE(1) DATA=
(S)</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0p=
x;"><span style=3D"font-size:10.0pt">// # of events - DATA(4)</span><br></p=
><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=
=3D"font-size:10.0pt">// Other - DATA(4)</span><br></p><p class=3D"" style=
=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">=
&nbsp;</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom:=
 0px;"><u><span style=3D"font-size:10.0pt">EVENTS</span></u><br></p><p clas=
s=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-s=
ize:10.0pt">// CEventBarLinear or
CEventBarPulse - SIZE(1) 00 DATA(S)</span><br></p><p class=3D"" style=3D"ma=
rgin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">// Eve=
nts - TRAK(2) STRT(2)
0000 ENDT(2) 0000 SPOS(2) 0000 EPOS(2) 0000 SIZE(1) TEXT(S) DATA(12 or 16) =
XX
XX</span><br></p><p style=3D"text-indent: 0.5in; margin-top: 0px; margin-bo=
ttom: 0px;" class=3D""><span style=3D"font-size:10.0pt">// DATA(12)
=3D Linear</span><br></p><p style=3D"text-indent: 0.5in; margin-top: 0px; m=
argin-bottom: 0px;" class=3D""><span style=3D"font-size:10.0pt">//
DATA(16) =3D Pulse</span><br></p><p class=3D"" style=3D"margin-top: 0px; ma=
rgin-bottom: 0px;"><span style=3D"font-size:10.0pt"><span style=3D"mso-spac=
erun:yes">&nbsp;&nbsp;&nbsp; </span><span style=3D"mso-tab-count:1">&nbsp;&=
nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; </span>//
XX XX decipher</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin=
-bottom: 0px;"><span style=3D"font-size:10.0pt"><span style=3D"mso-tab-coun=
t:
1">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;=
&nbsp;&nbsp;&nbsp; </span>// FF FF =3D New Event Type</span><br></p><p styl=
e=3D"margin-left: 0.5in; text-indent: 0.5in; margin-top: 0px; margin-bottom=
: 0px;" class=3D""><span style=3D"font-size:10.0pt">//followed by CEventBar=
Linear or CEventBarPulse -
SIZE(1) 00 DATA(S)</span><br></p><p style=3D"text-indent: 0.5in; margin-top=
: 0px; margin-bottom: 0px;" class=3D""><span style=3D"font-size:10.0pt">//
01 80 =3D Same as first event type</span><br></p><p style=3D"text-indent: 0=
.5in; margin-top: 0px; margin-bottom: 0px;" class=3D""><span style=3D"font-=
size:10.0pt">//
30 87 =3D Other event type</span><br></p><p style=3D"text-indent: 0.5in; ma=
rgin-top: 0px; margin-bottom: 0px;" class=3D""><span style=3D"font-size:10.=
0pt">//
00 00 =3D Last event</span><br></p><p class=3D"" style=3D"margin-top: 0px; =
margin-bottom: 0px;"><span style=3D"font-size:10.0pt">&nbsp;</span><br></p>=
<p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><u><span style=
=3D"font-size:10.0pt">AUDIO FILES</span></u><br></p><p class=3D"" style=3D"=
margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">//# =
of Files - DATA(2) 0000 (IF
FILES)</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom:=
 0px;"><span style=3D"font-size:10.0pt">// Audio File - SIZE(1)
FILE(S) SIZE(1) AUDIO_DEVICE(S)</span><br></p><p class=3D"" style=3D"margin=
-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">&nbsp;</sp=
an><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><u>=
<span style=3D"font-size:10.0pt">VIDEO FILES</span></u><br></p><p class=3D"=
" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:1=
0.0pt">// # of Files - DATA(2)
0000(IF FILES)</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin=
-bottom: 0px;"><span style=3D"font-size:10.0pt">// Video File - SIZE(1)
FILE(S) SIZE(1) AUDIO(S) SIZE(1) MONITOR(S) FULLSCREEN(1) XOFFSET(2) 0000
YOFFSET(2)</span><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bot=
tom: 0px;"><span style=3D"font-size:10.0pt">&nbsp;</span><br></p><p class=
=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-si=
ze:10.0pt">// Other - DATA(14)</span><br></p><p class=3D"" style=3D"margin-=
top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">&nbsp;</spa=
n><br></p><p class=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><u><=
span style=3D"font-size:10.0pt">TRACK SETTINGS</span></u><br></p><p class=
=3D"" style=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-si=
ze:10.0pt">// # of Tracks - DATA(2) 0000</span><br></p><p class=3D"" style=
=3D"margin-top: 0px; margin-bottom: 0px;"><span style=3D"font-size:10.0pt">=
// Tracks - SIZE(1) TEXT(S)
ADDR(2) 0000 CNTR(1) DATA(11) +VAL(2) 0000 -VAL(2) 0000 NEUT(2) 0000 ENBL(1=
) 00
0000 SIZE(1) PORT(S) DATA(12)</span><br></p><div><br></div><div>-----------=
------------------------<br></div><div><br></div><div>Thanks,<br></div><div=
>Nelson</div><div id=3D"Zm-_Id_-Sgn" data-sigid=3D"5979551000000009003" dat=
a-zbluepencil-ignore=3D"true"><div><br></div></div><div><br></div><div clas=
s=3D"zmail_extra_hr" style=3D"border-top: 1px solid rgb(204, 204, 204); hei=
ght: 0px; margin-top: 10px; margin-bottom: 10px; line-height: 0px;"><br></d=
iv><div class=3D"zmail_extra" data-zbluepencil-ignore=3D"true"><div><br></d=
iv><div id=3D"Zm-_Id_-Sgn1">---- On Wed, 27 Dec 2023 17:48:37 -0500 <b>Tom =
Mornini &lt;tmornini@me.com&gt;</b> wrote ---<br></div><div><br></div><bloc=
kquote id=3D"blockquote_zmail" style=3D"margin: 0px;"><div dir=3D"auto"><di=
v>Hey, thanks for getting back to me.<br></div><div><br></div><div>Yes, I=
=E2=80=99d love to know what you know about the format!<br></div><div><br><=
/div><div dir=3D"ltr"><div>-- Tom Mornini<br></div><div><div>-- Sent from i=
Phone<br></div><div>-- Forgive brevity and typos<br></div></div></div><div =
dir=3D"ltr"><div><br></div><blockquote>On Dec 27, 2023, at 2:03=E2=80=AFPM,=
 Sales &lt;<a href=3D"mailto:sales@lifeape.com" target=3D"_blank">sales@lif=
eape.com</a>&gt; wrote:<br><br></blockquote></div><blockquote><div dir=3D"l=
tr"><div>=EF=BB=BFHi Tom,<br></div><div><br></div><div>Thought I had replie=
d to this but maybe not. I know Jerry very well, he and I work together on =
a number of projects. I did reverse engineer the VSA format and am willing =
to share it if you still need it.&nbsp;<br></div><div><div><br></div><div d=
ir=3D"ltr"><div>Thanks,<br></div><div>Nelson<br></div></div><div dir=3D"ltr=
"><div><br></div><blockquote>On Nov 3, 2023, at 4:41=E2=80=AFPM, Tom Mornin=
i &lt;<a href=3D"mailto:tmornini@me.com" target=3D"_blank">tmornini@me.com<=
/a>&gt; wrote:<br><br></blockquote></div><blockquote><div dir=3D"ltr"><div>=
=EF=BB=BFSure.<br></div><div><br></div><div>My buddy runs SkullTronix and I=
=E2=80=99m helping him build a small controller board to operate his animat=
ronic products without the need to be tethered to a computer.<br></div><div=
><br></div><div>I would love to open source the documentation, and provide =
a Go language implementation, but that=E2=80=99s subject to your approval, =
of course.<br></div><div><br></div><div><div>I very much appreciate your co=
nsideration, this would save me so much time and hassle! =F0=9F=99=8F<br></=
div><div><br></div><div dir=3D"ltr"><div>-- Tom Mornini<br></div><div><div>=
-- Sent from iPhone<br></div><div>-- Forgive brevity and typos<br></div></d=
iv></div><div dir=3D"ltr"><div><br></div><blockquote>On Nov 3, 2023, at 4:2=
9 AM, Sales &lt;<a href=3D"mailto:sales@lifeape.com" target=3D"_blank">sale=
s@lifeape.com</a>&gt; wrote:<br><br></blockquote></div><blockquote><div dir=
=3D"ltr"><div>=EF=BB=BF<br></div><div style=3D"font-family :  Verdana,  Ari=
al,  Helvetica,  sans-serif; font-size :  10pt;"><div>Hi Tom,<br></div><div=
><br></div><div>I was able to determine the VSA file format and would be wi=
lling to share it. Can I ask for a bit more information about the project?<=
br></div><div><br></div><div>Thanks,<br></div><div>Nelson<br></div><div><br=
></div><div id=3D"x_-1584077653Zm-_Id_-Sgn"><div><br></div></div><div><br><=
/div><div class=3D"x_-1584077653zmail_extra_hr" style=3D"border-top :  1px =
solid rgb(204, 204, 204); min-height:  0px; margin-top :  10px; margin-bott=
om :  10px; line-height :  0px;"><br></div><div class=3D"x_-1584077653zmail=
_extra"><div><br></div><div id=3D"x_-1584077653Zm-_Id_-Sgn1">---- On Tue, 3=
1 Oct 2023 16:33:56 -0400 <b>Sales @ LifeApe &lt;<a href=3D"mailto:sales@li=
feape.com" target=3D"_blank">sales@lifeape.com</a>&gt;</b> wrote ---<br></d=
iv><div><br></div><blockquote id=3D"x_-1584077653blockquote_zmail" style=3D=
"margin :  0px;"><div><p>From: Tom Mornini<br> Email: <a href=3D"mailto:tmo=
rnini@me.com" target=3D"_blank">tmornini@me.com</a><br> Subject: Did you re=
verse engineer the VSA format?</p><p>Message Body:<br> Hello there.</p><p>I=
'm working on a software project and the person I'm building it for would l=
ove it to be able to read VSA file formats directly.<br></p><p>Do you have =
documentation of that format that you'd be willing to share?<br></p><p>If n=
ot, would you be willing to sell it?<br></p><p>I appreciate your time, I'd =
really rather not have to figure this out myself. <img class=3D"x_123951771=
wp-smiley" style=3D"height  :  1em; max-height  :  1em;" height=3D"1em" src=
=3D""><br></p><p>--<br> This e-mail was sent from a contact form on LifeApe=
 (<a href=3D"https://lifeape.com" target=3D"_blank">https://lifeape.com</a>=
)</p></div></blockquote></div><div><br></div></div><div><br></div></div></b=
lockquote></div></div></blockquote></div></div></blockquote></div></blockqu=
ote></div><div><br></div></div><br></body></html>
------=_Part_1160233_750255476.1704221568855--
```
