{
  description = "Dev environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {
        inherit system;
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      devShells = forAllSystems ({ system, pkgs }: {
        default = pkgs.mkShell {
          nativeBuildInputs = with pkgs; [
            go_1_22
            nodejs_21
          ];

          shellHook = ''
            echo "Development environment start.."
            echo "Installing lib dependency..."
            echo "air: "
            go install github.com/air-verse/air@latest
            echo "templ: "
            go install github.com/a-h/templ/cmd/templ@latest
            echo "To run development server, simply execute: "
            echo "make run"
          '';
        };
      });
    };
}