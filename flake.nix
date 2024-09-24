{
  description = "A basic browser built for browserjam001";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: {
    devShell.x86_64-linux = nixpkgs.legacyPackages.x86_64-linux.mkShell {
      buildInputs = with nixpkgs.legacyPackages.x86_64-linux; [
        libGL
        xorg.libXi
        xorg.libXcursor
        xorg.libXrandr
        xorg.libXinerama
        wayland
        libxkbcommon
      ];

    };
  };
}
