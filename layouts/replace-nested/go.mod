module go-my-playground/layouts/replace-nested

go 1.23.2

replace (
github.com/some-prominent-account/some-popular-module v1.2.3 => ./contrib/nested
)


require (
    github.com/some-prominent-account/some-popular-module v1.2.3
)