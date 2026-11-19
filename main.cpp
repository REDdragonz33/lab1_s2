#include "stack.h"
#include "queue.h"
#include "slist.h"
#include "dlist.h"
#include "array.h"
#include "tree.h"
#include "istructure.h"

#include <iostream>
#include <fstream>
#include <sstream>
#include <cctype>
#include <string>
#include <vector>
#include <map>

using namespace std;

void printHelp(const string& type); 

struct Registry {
    map<string, IStructure*> structures;
    
    ~Registry() {
        for (auto& p : structures) {
            delete p.second;
        }
    }
};

IStructure* createStructure(char type) {
    switch (type) {
        case 'S': return new Stack();
        case 'Q': return new Queue();
        case 'F': return new SList();
        case 'L': return new DList();
        case 'M': return new Array();
        case 'T': return new CompleteTree();
        default: return nullptr;
    }
}
// разбиваем строки на слова
vector<string> tokenize(const string& line) {
    vector<string> tokens;
    istringstream iss(line);
    string token;
    while (iss >> token) tokens.push_back(token);
    return tokens;
}
// выполнение 1 команды
void executeCommand(Registry& reg, const string& query) {
    vector<string> tokens = tokenize(query);
    if (tokens.empty()) return;
    
    string cmd = tokens[0];

    if (cmd == "HELP") {
        if (tokens.size() > 1) {
            string target = tokens[1];
            for (char& c : target) c = toupper(c);
            printHelp(target);
        } else {
            printHelp("");
        }
        return;
    }

    if (cmd == "PRINT") {
        if (tokens.size() < 2) { cout << "PRINT: missing name" << endl; return; }
        string name = tokens[1];
        if (reg.structures.count(name)) {
            reg.structures[name]->print();
        } else {
            cout << "Structure '" << name << "' not found" << endl;
        }
        return;
    }

    if (tokens.size() < 2) { cout << "Missing structure name" << endl; return; }
    string name = tokens[1];
    char type = cmd[0];
    
    // Если структуры нет, создаем ее
    if (!reg.structures.count(name)) {
        IStructure* newStruct = createStructure(type);
        if (!newStruct) {
            cout << "Unknown structure type: " << type << endl;
            return;
        }
        reg.structures[name] = newStruct;
    }

    reg.structures[name]->execute(cmd, tokens);
}
void saveToFile(const Registry& reg, const string& filename) {
    ofstream out(filename);
    if (!out) {
        cout << "Error: Cannot open file for writing" << endl;
        return;
    }

    for (const auto& pair : reg.structures) {
        const string& name = pair.first;
        IStructure* structPtr = pair.second;
        out << structPtr->getType() << " " << name << " ";
        structPtr->save(out); 
        out << endl;
    }
    cout << "Saved to " << filename << endl;
}

void loadFromFile(Registry& reg, const string& filename) {
    ifstream in(filename);
    if (!in) {
        cout << "Error: File not found" << endl;
        return;
    }
    
    string line;
    while (getline(in, line)) {
        if (line.empty()) continue;
        
        istringstream iss(line);
        char type;
        string name;

        if (!(iss >> type >> name)) continue;

        if (!reg.structures.count(name)) {
            IStructure* newStruct = createStructure(type);
            if (newStruct) {
                reg.structures[name] = newStruct;
            } else {
                cout << "Warning: Unknown structure type '" << type << "'" << endl;
                continue;
            }
        }
        reg.structures[name]->load(iss); 
    }
    cout << "Loaded from " << filename << endl;
}

void printHelp(const string& type) {
    if (type == "STACK" || type == "S" || type == "SPUSH" || type == "SPOP") {
        cout << "=== СТЕК (Stack) ===" << endl;
        cout << "SPUSH <имя> <значение>  - Добавить элемент на вершину" << endl;
        cout << "SPOP <имя>              - Удалить элемент с вершины" << endl;
        cout << "PRINT <имя>             - Показать стек" << endl;
    } 
    else if (type == "QUEUE" || type == "Q" || type == "QPUSH" || type == "QPOP") {
        cout << "=== ОЧЕРЕДЬ (Queue) ===" << endl;
        cout << "QPUSH <имя> <значение>  - Добавить элемент в конец" << endl;
        cout << "QPOP <имя>              - Удалить элемент из начала" << endl;
        cout << "PRINT <имя>             - Показать очередь" << endl;
    }
    else if (type == "ARRAY" || type == "M" || type == "MPUSH") {
        cout << "=== МАССИВ (Array) ===" << endl;
        cout << "MPUSH <имя> <значение>  - Добавить в конец" << endl;
        cout << "MPUSH <имя> <индекс> <значение> - Вставить по индексу" << endl;
        cout << "MDEL <имя> <индекс>     - Удалить по индексу" << endl;
        cout << "MGET <имя> <индекс>     - Получить по индексу" << endl;
        cout << "MSET <имя> <индекс> <значение> - Заменить по индексу" << endl;
        cout << "MLEN <имя>              - Узнать длину" << endl;
        cout << "PRINT <имя>             - Показать массив" << endl;
    }
    else if (type == "SLIST" || type == "F" || type == "FPUSH") {
        cout << "=== ОДНОСВЯЗНЫЙ СПИСОК (SList) ===" << endl;
        cout << "FPUSH <имя> HEAD <значение>   - Вставить в голову" << endl;
        cout << "FPUSH <имя> TAIL <значение>   - Вставить в хвост" << endl;
        cout << "FPUSH <имя> AFTER <цель> <значение> - Вставить после" << endl;
        cout << "FPUSH <имя> BEFORE <цель> <значение> - Вставить перед" << endl;
        cout << "FDEL <имя> <значение>         - Удалить по значению" << endl;
        cout << "FDEL <имя> AFTER <цель>       - Удалить ПОСЛЕ цели" << endl;
        cout << "FDEL <имя> BEFORE <цель>      - Удалить ПЕРЕД целью" << endl;
        cout << "PRINT <имя>                   - Показать список" << endl;
    }
    else if (type == "DLIST" || type == "L" || type == "LPUSH") {
        cout << "=== ДВУСВЯЗНЫЙ СПИСОК (DList) ===" << endl;
        cout << "LPUSH <имя> HEAD <значение>   - Вставить в голову" << endl;
        cout << "LPUSH <имя> TAIL <значение>   - Вставить в хвост" << endl;
        cout << "LPUSH <имя> AFTER <цель> <значение> - Вставить после" << endl;
        cout << "LPUSH <имя> BEFORE <цель> <значение> - Вставить перед" << endl;
        cout << "LDEL <имя> <значение>         - Удалить по значению" << endl;
        cout << "PRINT <имя>                   - Показать список" << endl;
        cout << "LDEL <имя> AFTER <цель>       - Удалить ПОСЛЕ цели" << endl;
        cout << "LDEL <имя> BEFORE <цель>      - Удалить ПЕРЕД целью" << endl;
    }
    else if (type == "TREE" || type == "T" || type == "TINSERT") {
        cout << "=== ДЕРЕВО (Complete Tree) ===" << endl;
        cout << "TINSERT <имя> <число>   - Добавить число" << endl;
        cout << "TGET <имя> <число>      - Поиск числа (TRUE/FALSE)" << endl;
        cout << "TCHECK <имя>            - Проверить на полноту (TRUE/FALSE)" << endl;
        cout << "PRINT <имя>             - Показать дерево" << endl;
    }
    else {
        cout << "=== ДОСТУПНЫЕ СТРУКТУРЫ ===" << endl;
        cout << "HELP STACK   - Справка по стеку" << endl;
        cout << "HELP QUEUE   - Справка по очереди" << endl;
        cout << "HELP ARRAY   - Справка по массиву" << endl;
        cout << "HELP SLIST   - Справка по односвязному списку" << endl;
        cout << "HELP DLIST   - Справка по двусвязному списку" << endl;
        cout << "HELP TREE    - Справка по дереву" << endl;
        cout << "PRINT <имя>  - Вывести любую структуру" << endl;
    }
}

int main(int argc, char** argv) {
    string filename;
    string query;
    
    for (int i = 1; i < argc; ++i) {
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc) filename = argv[++i];
        else if (arg == "--query" && i + 1 < argc) query = argv[++i];
    }
    
    Registry reg;
    
    if (!filename.empty()) loadFromFile(reg, filename);
    
    if (!query.empty()) {
        executeCommand(reg, query);
    } else {
        cout << "DBMS ready. Type EXIT to quit." << endl;
        string line;
        while (true) {
            cout << "> ";
            if (!getline(cin, line) || line == "EXIT") break;
            if (!line.empty()) executeCommand(reg, line);
        }
    }
    
    if (!filename.empty()) saveToFile(reg, filename);
    
    return 0;
}