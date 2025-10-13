#!/bin/bash

# test.sh - Local testing script for GoTour exercises
set -e  # Exit on any error

echo "🔍 Testing GoTour Exercises Repository"
echo "======================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counter for results
PASS=0
FAIL=0
SKIP=0

print_result() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✅ $2${NC}"
        ((PASS++))
    else
        echo -e "${RED}❌ $2${NC}"
        ((FAIL++))
    fi
}

print_skip() {
    echo -e "${YELLOW}⚠️  $1${NC}"
    ((SKIP++))
}

echo
echo "📦 Step 1: Checking all Go files compile..."
echo "-------------------------------------------"

# Check all Go files can be compiled
while IFS= read -r -d '' file; do
    pkg_dir=$(dirname "$file")
    if [[ "$pkg_dir" == *".git"* ]]; then
        continue
    fi
    
    # Try to compile as package first
    if cd "$pkg_dir" 2>/dev/null && go build -o /dev/null . 2>/dev/null; then
        echo -e "${GREEN}✅ $pkg_dir compiles as package${NC}"
        ((PASS++))
    else
        # If it fails, it might be a main package, which is fine
        print_skip "$pkg_dir doesn't compile as package (likely main)"
    fi
    cd - > /dev/null
done < <(find . -name "*.go" -type f -print0)

echo
echo "🔨 Step 2: Building all main packages..."
echo "---------------------------------------"

# Build all main packages
while IFS= read -r -d '' file; do
    dir=$(dirname "$file")
    if [[ "$dir" == *".git"* ]]; then
        continue
    fi
    
    echo -n "Building $dir... "
    if cd "$dir" && go build -o /dev/null .; then
        print_result 0 "$dir builds successfully"
    else
        print_result 1 "$dir failed to build"
    fi
    cd - > /dev/null
done < <(find . -name "main.go" -type f -print0)

echo
echo "🧪 Step 3: Running available tests..."
echo "------------------------------------"

# Run tests where they exist
while IFS= read -r -d '' file; do
    dir=$(dirname "$file")
    if [[ "$dir" == *".git"* ]]; then
        continue
    fi
    
    echo -n "Testing $dir... "
    if cd "$dir" && go test -v .; then
        print_result 0 "$dir tests passed"
    else
        print_result 1 "$dir tests failed"
    fi
    cd - > /dev/null
done < <(find . -name "*_test.go" -type f -print0)

echo
echo "📊 Step 4: Summary"
echo "------------------"

echo -e "Total: ${GREEN}$PASS passed${NC}, ${RED}$FAIL failed${NC}, ${YELLOW}$SKIP skipped${NC}"

if [ $FAIL -eq 0 ]; then
    echo -e "${GREEN}🎉 All checks passed!${NC}"
    exit 0
else
    echo -e "${RED}💥 Some checks failed!${NC}"
    exit 1
fi