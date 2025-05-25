// GO-Minus Windows Runtime Library
// Bu dosya Windows'ta printf ve diğer C runtime fonksiyonlarını sağlar

#include <windows.h>
#include <stdio.h>

// Printf implementation using Windows API
int printf(const char* format, ...) {
    va_list args;
    va_start(args, format);
    
    char buffer[1024];
    int result = vsnprintf(buffer, sizeof(buffer), format, args);
    
    HANDLE hConsole = GetStdHandle(STD_OUTPUT_HANDLE);
    DWORD written;
    WriteConsoleA(hConsole, buffer, strlen(buffer), &written, NULL);
    
    va_end(args);
    return result;
}

// Entry point for Windows
int main(void);

int WINAPI WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance, LPSTR lpCmdLine, int nCmdShow) {
    // Allocate console for output
    AllocConsole();
    freopen("CONOUT$", "w", stdout);
    
    // Call user's main function
    return main();
}

// Alternative entry point
int mainCRTStartup(void) {
    return main();
}
