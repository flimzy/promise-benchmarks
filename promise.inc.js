$global.EmptyPromise = function() {
    return new Promise(function(resolve, reject) {
        resolve('Success!');
    });
};

$global.EmptyCallback = function(resolve, reject) {
    resolve('Success!');
};
