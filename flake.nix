{
  description = "Flake for unofficial TUI client for music.yandex.ru";

  inputs.nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
  in {
    devShells.x86_64-linux.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        go-task
      ];
      shellHook = ''
        go version
        task --version
      '';
    };
  };
}
