# VSA file format support for Go

This package provides support for reading VSA files.

## History

This code based upon documentation provided by Nelson Bairos of [LifeApe](https://lifeape.com)

## Format in bytes

All numbers are encoded in (sign or unsigned?) little endian format.

Strings are stored as a number of bytes of length followed by that many bytes of the string.

They are denoted as (bytes of length, S).

Fields named unknown are not well understood by the author.

## Unknowns

### unknownOne

Nelson documented this as Version.

12 bytes is a gigantic number of bytes for a version.

In my 5 test files from a single user the first 12 bytes were always `0ad7a3703d0a184001000000`

Doesn't look like a string. Too big for a number. Is anything else 12 bytes long?

What does it mean? 🤷🏻‍♂️

Perhaps it shows up in pi if you go our far enough? 🤣

Google comes up empty all the way down to `0ad7a3703d0a18` and `0ad7a3703d0a` (removing from the end one byte at a time) which both returned this [URL](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwjyityd95WEAxUiszEKHWANDlkQFnoECBEQAQ&url=https%3A%2F%2Fwww.faulknerhondadoylestown.com%2Forder-parts%2F%3Fpath%3D%252Foem-parts%252Fhonda-lever-sub-assy-change-54100sdaa01%253Fc%253DbT0xJmw9NCZuPVJlY29tbWVuZGVkIFByb2R1Y3Rz&usg=AOvVaw3upsmAatIKO_9bxMOICk23&opi=89978449) (and no, it doesn't exist in the URL or in the response, so WTF?) after which `0ad7a3703d` again comes up empty. Which struck me as bizarre enough to earn it's place in this document.

Oh, interesting! `0ad7a3` came up with links to a color hex code site. Does the VSA file creator program happen to have 4 configurable colors? 🤔 That sounds fairly plausible.

Help? 🙏🏻 @tmornini

### unknownTwo

Nelson documented as other, which I believe indicated that he didn't know what it was. Having already changed `version` to `unkown` I chose rename both to `unknownOne` and `unknownTwo` to indicate that I don't know what they are.

### unknownFive

Nelson didn't document this field. Since it's between events, I've chosen to discard it until I have a better understanding of what it is.

### unknownSix

Nelson documented this field. Since it's between sections, I've chosen to discard it until I have a better understanding of what it is.

### unknownEight and unknownNine

Nelson documented these fields as `00 0000` so I've kept them as such until more is known.

### unknownTen

Nelson documented as `DATA(12)` so I've kept it as such until more is known.

``` text
HEADER
    unknownOne(12)
    LEVEL(1,S)
    OPTIONS(1,S)
    EMAIL(1,S)
    EVENTCOUNT(4)
    unknownTwo(4)
    DEFAULTEVENTTYPE(2,S)

EVENTS
    EVENT
        _type
            "CEventBarLinear" or "CEventBarPulse"
        TRACK(2)
        STARTTIME(4)
        ENDTIME(4)
        STARPOSITION(4)
        ENDPOSITION(4)
        TEXT(1,S)
        case _Type
            "CEventBarLinear"
                DATA(12)
            "CEventBarPulse"
                DATA(16) Pulse
        NEXTEVENTTYPE(2)
            00 00 = documented by Nelson as "Last event"
            01 00 = Final event
            01 80 = Default event type
            30 87 = Other event type (not Default event type?)
            FF FF = New Event Type
                unknownFive(2)
                currentEventType(2,S)

AUDIO FILES
    COUNT(4)
        AUDIO FILE
            FILE(1,S)
            AUDIO_DEVICE(1,S)

VIDEO FILES
    COUNT(4)
        VIDEO FILE
            FILE(1,S)
            AUDIO(1,S)
            MONITOR(1,S)
            FULLSCREEN(1)
            XOFFSET(4)
            YOFFSET(4)

unknownSix(12)

TRACK SETTINGS
    COUNT(4)
        TRACK SETTING
            TEXT(1,S)
            ADDR(4)
            CNTR(1)
            DATA(11)
            +VAL(4)
            -VAL(4)
            NEUT(4)
            ENBL(1)
            unknownEight(1)
            unknownNine(2)
            PORT(1,S)
            unknownTen(12)
```
