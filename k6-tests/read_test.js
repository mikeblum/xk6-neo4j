import { check } from 'k6';
import neo4j from 'k6/x/neo4j';

const db = neo4j.openWithConf();

export function setup() { }

export function teardown() {  
    db.close();
}

export default function () {
    check(db.return(), {
        // why doesn't === work here returning int64?
        'is response OK': (r) => r == 1,
    });
}
