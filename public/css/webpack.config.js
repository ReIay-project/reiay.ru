const path = require('path');

module.exports = {
    entry: './public/js/scripts.js',
    output: {
        filename: 'build.js',
        path: path.resolve(__dirname, 'public/js')
    },
    mode: 'production',
    module: {
        rules: [
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
            }
        ]
    }
};
