#!/usr/bin/env bash
set -e

OS="$(uname -s)"
ARCH="$(uname -m)"

case "${OS}" in
    Linux*)     OS_NAME="linux";;
    Darwin*)    OS_NAME="darwin";;
    *)          echo "Unsupported OS: ${OS}"; exit 1;;
esac

case "${ARCH}" in
    x86_64|amd64) ARCH_NAME="amd64";;
    arm64|aarch64) ARCH_NAME="arm64";;
    *)            echo "Unsupported architecture: ${ARCH}"; exit 1;;
esac

REPO="Thedtk24/menv"
echo "Searching for the latest release of ${REPO}..."

LATEST_TAG=$(curl -s "https://api.github.com/repos/${REPO}/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "${LATEST_TAG}" ] || [ "${LATEST_TAG}" = "null" ]; then
    echo "Error retrieving the latest tag. Ensure the repository exists and has a release with binaries."
    exit 1
fi

echo "Latest release: ${LATEST_TAG}"

TARBALL="menv_${OS_NAME}_${ARCH_NAME}.tar.gz"
DOWNLOAD_URL="https://github.com/${REPO}/releases/download/${LATEST_TAG}/${TARBALL}"

TMP_DIR=$(mktemp -d)
trap 'rm -rf -- "${TMP_DIR}"' EXIT

cd "${TMP_DIR}"
echo "📥 Downloading from ${DOWNLOAD_URL}..."
curl -sL --fail -o "${TARBALL}" "${DOWNLOAD_URL}" || (echo "Download failed"; exit 1)

echo "Extracting..."
tar -xzf "${TARBALL}"

INSTALL_DIR="${HOME}/.local/bin"
mkdir -p "${INSTALL_DIR}"

echo "Installing to ${INSTALL_DIR}/menv..."
mv menv "${INSTALL_DIR}/menv"
chmod +x "${INSTALL_DIR}/menv"

echo "✓ Installation successful!"
echo "Make sure ${INSTALL_DIR} is in your PATH."
