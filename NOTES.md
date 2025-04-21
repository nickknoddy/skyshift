## Points

1. Check out other libraries for image transformations
1. Libraries which run on GPU is preferrable
1. `gocv` library needs system level dependencies
1. Benchmark all the transformations and develop a table so it can be tracked
1. `opencv` python sharpen code is much faster
1. Develop a small prototype in cuda
1. Keep the output image format same as input
1. Add functionality for chaining transformation
1. Transformation query paramater can be replace with single query enclosing all transformations

   - w: width, h: height, sh: sharpen, br: brightness
   - query: w-360,h-360,sh-5,br-5
   - query string can be parsed
   - parsing the string will have some latency which will add to overall api response
   - sanitization of query is also important

1. Enable pre-transformations while uploading images which can be cached
   - will be helpful in cases of resizing, crop and compression
1. In chaining functionality, use buffers if possible between transformations and only convert it the image format at the last transformation
