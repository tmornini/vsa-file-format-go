# VSA file format support for Go

This package provides support for reading VSA files.

## History

This code based upon documentation provided by Nelson (FIXME) of [LifeApe](https://lifeape.com)

## Format

``` text
HEADER
    unknownOne(12)
    SIZE(1) LEVEL(S)
    SIZE(1) OPTIONS(S)
    SIZE(1) EMAIL(S)
    EVENTCOUNT(4)
    unknownTwo(4)

EVENTS
    SIZE(2) DEFAULTEVENTTYPE(S)
        EVENT
            TRACK(2)
            STARTTIME(4)
            ENDTIME(4)
            STARPOSITION(4)
            ENDPOSITION(4)
            SIZE(1) KIND(S)            <--\
                "CEventBarLinear"         |
                    DATA(12)              | <-\
                "CEventBarPulse"          |   | or
                    DATA(16) Pulse        | <-|
            NEXTEVENTTYPE(2)              |   |
                FF FF = New Event Type >--/   |
                01 80 = Default event type >--/
                30 87 = Other event type ? perhaps observed but not understood ?
                00 00 = Last event

AUDIO FILES
    COUNT(4)
        SIZE(1) FILE(S)
        SIZE(1) AUDIO_DEVICE(S)

VIDEO FILES
    COUNT(4)
        SIZE(1) FILE(S)
        SIZE(1) AUDIO(S)
        SIZE(1) MONITOR(S)
        FULLSCREEN(1)
        XOFFSET(4)
        YOFFSET(4)

unknown3(12)

TRACK SETTINGS
    COUNT(4)
        SIZE(1) TEXT(S)
        ADDR(4)
        CNTR(1)
        unknown4(11)
        +VAL(4)
        -VAL(4)
        NEUT(4)
        ENBL(1)
        unknown5(1)
        unknown6(2)
        SIZE(1) PORT(S)
        unknown7(12)
```
