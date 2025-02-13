let gulp       = require('gulp'),
    livereload = require('gulp-livereload');

// gulp LiveReload
gulp.task('LiveReload', function () {
    livereload.listen();
    gulp.watch('./**/*', {delay: 300}, async function () {
        livereload.reload();
    });
});