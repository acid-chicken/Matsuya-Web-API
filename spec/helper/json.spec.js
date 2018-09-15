import {JsonUtil} from '../../helper/json';

describe("JSON Utility", function() {
    it("Is menu observer not returned undefined", function() {
      const jsonUtil = new JsonUtil();
      expect(jsonUtil.getAllMenu()).not.toBe(undefined);
    });
});
