(function() {
  "use strict";

  angular.module("labApp").controller("PlayersController", PlayersController);

  /* @ngInject */
  function PlayersController($state, playersService) {
    var vm = this;

    vm.players = [];

    activate();

    vm.player = {
      Name: "",
      ID: "",
      LastName: ""
    };

    vm.createPlayer = createPlayer;
    vm.deletePlayer = deletePlayer;
    vm.deletePlayers = deletePlayers;

    //////////////////

    function activate() {
      getPlayers();
    }

    function createPlayer() {
      playersService
        .create(vm.player)
        .then(function() {
          $state.reload();
        })
        .catch(function(error) {});
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

    function deletePlayers(playerId) {
      console.log(playerId);
      playersService
        .eliminateAll()
        .then(function(response) {
          $state.reload();
        })
        .catch(function(error) {});
    }
  }
})();
