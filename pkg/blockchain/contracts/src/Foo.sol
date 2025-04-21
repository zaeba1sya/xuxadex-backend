// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Foo {
    struct FooStruct {
        address foo1;
        uint256 foo2;
        string foo3;
    }

    FooStruct public localFoo;

    event FooChangeEvent(
        address sender,
        address foo1,
        uint256 foo2,
        string foo3
    );

    constructor() {
        localFoo = FooStruct(msg.sender, 123, "Hello, world");
    }

    function setAddress(address _foo1) external {
        localFoo.foo1 = _foo1;
        emit FooChangeEvent(
            msg.sender,
            localFoo.foo1,
            localFoo.foo2,
            localFoo.foo3
        );
    }

    function setUint(uint256 _foo2) external {
        localFoo.foo2 = _foo2;
        emit FooChangeEvent(
            msg.sender,
            localFoo.foo1,
            localFoo.foo2,
            localFoo.foo3
        );
    }

    function setString(string calldata _foo3) external {
        localFoo.foo3 = _foo3;
        emit FooChangeEvent(
            msg.sender,
            localFoo.foo1,
            localFoo.foo2,
            localFoo.foo3
        );
    }
}
