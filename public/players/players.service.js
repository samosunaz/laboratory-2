(function() {
  "use strict";

  angular.module("labApp").service("playersService", playersService);

  playersService.$inject = ["$http", "$q"];
  function playersService($http, $q) {
    var apiUrl = "http://localhost:8000";

    this.all = all;
    this.create = create;
    this.eliminate = eliminate;
    this.eliminateAll = eliminateAll;

    ////////////////

    function onSuccess(response) {
      console.log(response);
      return response.data;
    }

    function onError(error) {
      console.log(error);
      return $q.reject(error);
    }

    function all() {
      return $http
        .get(`${apiUrl}/players`)
        .then(onSuccess)
        .catch(onError);
    }

    function create(player) {
      return $http
        .post(`${apiUrl}/players`, player)
        .then(onSuccess)
        .catch(onError);
    }

    function eliminate(playerId) {
      return $http
        .delete(`${apiUrl}/players/${playerId}`)
        .then(onSuccess)
        .catch(onError);
    }

    function eliminateAll() {
      return $http
        .delete(`${apiUrl}/players`)
        .then(onSuccess)
        .catch(onError);
    }
  }
})();
