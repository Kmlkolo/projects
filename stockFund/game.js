// metody uprzywilejowane, prywatne zmienne wspólne dla wszystkich instancji obiektu
// event handling

var Trader = (function() {
      var fundResult = 0;
      function InnerTrader() {
            this.result = result;
      }
      InnerTrader.prototype.balanceChange = function(amount) {
            InnerTrader.account += amount;
      };
      
}());
