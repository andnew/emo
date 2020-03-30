module ch16

go 1.13

require (
    ch16/con v0.0.0 //indirect
    ch16/ext v0.0.0 //indirect
)

replace (
    ch16/con v0.0.0 => ./ch16/con
    ch16/ext v0.0.0 => ./ch16/ext
)