#ifndef DLIST_H
#define DLIST_H

#include "istructure.h"
#include <iostream>
#include <string>
using namespace std;

struct DNode {
    string data;
    DNode* prev;
    DNode* next;
    DNode(string val, DNode* p = nullptr, DNode* n = nullptr);
};

class DList : public IStructure {
private:
    DNode* head;
    DNode* tail;

public:
    DList();
    ~DList();

    void insertHead(string val);
    void insertTail(string val);
    void insertAfter(string target, string val);
    void insertBefore(string target, string val);
    void remove(string val);
    
    void print() override;
    void execute(const string& cmd, const vector<string>& tokens) override;
    void save(ostream& out) const override;
    void load(istream& in) override;
    char getType() const override;
    void removeAfter(string target);
    void removeBefore(string target);
};

#endif