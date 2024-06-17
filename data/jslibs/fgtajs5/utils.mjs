export function generateUUID() {
    if (window.crypto && window.crypto.getRandomValues) {
        var buffer = new Uint16Array(8);
        window.crypto.getRandomValues(buffer);
        
        var hexValues = [];
        for (var i = 0; i < buffer.length; i++) {
            hexValues.push(buffer[i].toString(16));
        }
        
        return (hexValues[0] + hexValues[1] + "-" + hexValues[2] + "-" + hexValues[3] + "-" + hexValues[4] + "-" + hexValues[5] + hexValues[6] + hexValues[7]);
    } else {
        // Fallback untuk lingkungan yang tidak mendukung crypto.getRandomValues
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
            var r = Math.random() * 16 | 0,
                v = c === 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }
}
