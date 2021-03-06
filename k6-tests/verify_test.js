import { check } from 'k6';
import neo4j from 'k6/x/neo4j';

const db = neo4j.open("neo4j://localhost:7687", "neo4j", "localhost");

export function setup() { }

export function teardown() {  
    db.close();
}

export default function () {
    check(db.verify(), {
        'is status OK': (status) => status === true,
    });
}
