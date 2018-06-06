# imgn
batch jpeg editor on commandline

        imgn [output quantity] [file] [-1] [adjustments] [-n] [adj1,adj2]

##adjustments:

b       adjust brightness of image
    :: number between -1 (reduce brightness) and 1 (increase brightness)
    :: number greater than 0 increases brightness
        example:    -n b=0.204
                    -n b=-0.456

c       adjust contrast of image
    :: number between -1 (reduce contrast) and 3 (increase contrast)
    :: number greater than 0 increases contrast
        example:
                    -n c=1.55
                    -n c=-0.405

d       desaturate image
    ::  luminosity      0.21*red + 0.72*g + 0.07*b
    ::  average         (r + g + b) / 3
    ::  lightness       (max + min) / 2
        example:
                    -n d=luminosity
                    -n d=average
                    -n d=lightness

df      desaturate with custom formula
    ::  weight*red + weight*green + weight*blue
    ::  weights are user inputed as floats
        example:
                    -n df=[red]_[green]_[blue]
                    -n df=0.20_0.70_0.10