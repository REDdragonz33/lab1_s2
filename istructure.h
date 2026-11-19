#ifndef ISTRUCTURE_H
#define ISTRUCTURE_H

#include <iostream>
#include <string>
#include <vector>
using namespace std;

class IStructure {
public:
    virtual ~IStructure() {} 
    virtual void execute(const string& cmd, const vector<string>& tokens) = 0;
    virtual void print() = 0;
    virtual void save(ostream& out) const = 0;
    virtual void load(istream& in) = 0;
    virtual char getType() const = 0;
};

#endif