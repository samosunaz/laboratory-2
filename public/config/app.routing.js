(function() {
  "use strict";

  angular.module("labApp").config(config);

  function config($locationProvider, $stateProvider, $urlRouterProvider) {
    let homeState = {
      name: "home",
      url: "/",
      controller: "PlayersController",
      controllerAs: "vm",
      templateUrl: "./players/players.html"
    };

    $stateProvider.state(homeState);

    $urlRouterProvider.otherwise("/");

    $locationProvider.html5Mode({
      enabled: true,
      requireBase: true
    });
  }
})();
