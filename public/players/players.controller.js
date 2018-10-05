(function() {
  "use strict";

  angular.module("labApp").controller("PlayersController", PlayersController);

  /* @ngInject */
  function PlayersController($state, playersService) {
    var vm = this;

    vm.players = [];

    activate();

    vm.deletePlayer = deletePlayer;

    //////////////////

    function activate() {
      getPlayers();
    }

    function getPlayers() {
      playersService
        .all()
        .then(function(players) {
          vm.players = players;
        })
        .catch(function(error) {});
    }

    function deletePlayer(playerId) {
      console.log(playerId);
      playersService
        .eliminate(playerId)
        .then(function(response) {
          $state.reload();
        })
        .catch(function(error) {});
    }
  }
})();
