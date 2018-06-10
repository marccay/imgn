# imgn
batch jpeg editor on commandline

        imgn [output quantity] [file / folder] [--options] [-1] [adjustments] [-n] [adj1,adj2]

## options:

###### "--train"
- include all images inside nested folders to be processed
    '''
    example:
        imgn -4 /home/folder/ --train -1 d=lightness...
    '''

## adjustments:

###### "b=" adjust brightness of image
- number between -1 (reduce brightness) and 1 (increase brightness)
- number greater than 0 increases brightness
    ```
    example:
        -n b=0.204
        -n b=-0.456
    ```

###### "c=" adjust contrast of image
- number between -1 (reduce contrast) and 2 (increase contrast)
- number greater than 0 increases contrast
    ```
    example:
        -n c=1.55
        -n c=-0.405
    ```
###### "d=" desaturate image
- luminosity
    ```
    0.21 * red + 0.72 * green + 0.07 * blue
    ```
- average
    ```
    (r + g + b) / 3
    ```
- lightness
    ```
    (max + min) / 2
    ```
    ```
    example:
        -n d=luminosity
        -n d=average
        -n d=lightness
    ```

###### "df=" desaturate with custom formula
- weight * red + weight * green + weight * blue
- weights are user inputed as floats
    ```
    example:
        -n df=[red]_[green]_[blue]
        -n df=0.20_0.70_0.10
    ```

###### "h=" highlight recovery
- a number between 0 and 1
- a higher number (closer to 1) recovers more highlights
    ```
    example:
        -n h=0.502
    ```

###### "s=" shadow recovery
- a number between 0 and 1
- a higher number (closer to 1) recovers more shadow
    ```
    example:
        -n s=0.880
    ```

###### "bx=" border
- [border pixel size]_[r]_[g]_[b]
- r, b, b values between 0 and 255
    ```
    example:
        -n bx=200 (default black border)
        -n bx=200_255_255_255 (white border)
        -n bx=200_255_50_50     
    ```