#!/bin/sh


set -e

# Define standard terminal color codes for styled outputs
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo "${BLUE}===================================================${NC}"
echo "${BLUE}  S.T.E.V.E. Cart - Display Service Starting...    ${NC}"
echo "${BLUE}===================================================${NC}"

# 1. Print current environment details for debugging
echo "${BLUE}[1/5] Checking environment state...${NC}"
echo "Current working directory: $(pwd)"
echo "Container directory contents:"
ls -la

# 2. Check if running on Alpine Linux and install Node/NPM if missing
echo "${BLUE}[2/5] Verifying system dependencies...${NC}"
if [ -f /etc/alpine-release ]; then
    echo "Alpine Linux detected."
    if ! command -v node >/dev/null 2>&1 || ! command -v npm >/dev/null 2>&1; then
        echo "${YELLOW}Node.js or NPM is not installed. Installing system packages...${NC}"
        apk update
        apk add --no-cache nodejs npm
        echo "${GREEN}Node.js and NPM successfully installed via apk!${NC}"
    else
        echo "${GREEN}Node.js ($(node -v)) and NPM ($(npm -v)) are already installed.${NC}"
    fi
else
    echo "Non-Alpine OS environment."
    if ! command -v node >/dev/null 2>&1 || ! command -v npm >/dev/null 2>&1; then
        echo "${RED}Error: Node.js and NPM are required but not installed!${NC}"
        exit 1
    fi
fi

# 3. Verify project structure
echo "${BLUE}[3/5] Verifying project structure...${NC}"
if [ ! -f "package.json" ]; then
    echo "${RED}Error: package.json not found in $(pwd)!${NC}"
    exit 1
fi
echo "${GREEN}package.json found successfully.${NC}"

# 4. Install Node dependencies if needed
echo "${BLUE}[4/5] Checking node_modules and dependencies...${NC}"
if [ ! -d "node_modules" ]; then
    echo "${YELLOW}node_modules not found. Installing dependencies...${NC}"
    npm install
    echo "${GREEN}Dependencies installed successfully!${NC}"
else
    echo "${GREEN}node_modules folder already exists. Skipping full install. Run 'npm install' inside container if you need to update package.json packages.${NC}"
fi

# 5. Start the Vite development server
echo "${BLUE}[5/5] Launching Vite development server...${NC}"
echo "${YELLOW}Starting server on host 0.0.0.0 (port 5173)...${NC}"
exec npm run dev -- --host 0.0.0.0
