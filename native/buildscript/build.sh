cd ..

echo "[BUILDSCRIPT] Checking-out the source code..."
cd libsodium
#git checkout 1.0.12

echo "[BUILDSCRIPT] Running autogen.sh..."
bash autogen.sh

echo "[BUILDSCRIPT] Running configure script..."
bash configure --disable-shared --prefix=$PWD/..

echo "[BUILDSCRIPT] Running make..."
make -j4
make install

