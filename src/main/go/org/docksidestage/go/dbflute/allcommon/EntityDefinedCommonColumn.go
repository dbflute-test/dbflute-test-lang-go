/*
 * Copyright 2014-2015 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language
 * governing permissions and limitations under the License.
 */
package org.docksidestage.go.dbflute.allcommon;

/**
 * The interface of entity defined common columns.
 * @author DBFlute(AutoGenerator)
 */
public interface EntityDefinedCommonColumn extends Entity {

    /**
     * Get the value of registerDatetime.
     * @return The value of registerDatetime. (NullAllowed)
     */
    public java.time.LocalDateTime getRegisterDatetime();

    /**
     * Set the value of registerDatetime.
     * @param registerDatetime The value of registerDatetime. (NullAllowed)
     */
    public void setRegisterDatetime(java.time.LocalDateTime registerDatetime);

    /**
     * Get the value of registerUser.
     * @return The value of registerUser. (NullAllowed)
     */
    public String getRegisterUser();

    /**
     * Set the value of registerUser.
     * @param registerUser The value of registerUser. (NullAllowed)
     */
    public void setRegisterUser(String registerUser);

    /**
     * Get the value of updateDatetime.
     * @return The value of updateDatetime. (NullAllowed)
     */
    public java.time.LocalDateTime getUpdateDatetime();

    /**
     * Set the value of updateDatetime.
     * @param updateDatetime The value of updateDatetime. (NullAllowed)
     */
    public void setUpdateDatetime(java.time.LocalDateTime updateDatetime);

    /**
     * Get the value of updateUser.
     * @return The value of updateUser. (NullAllowed)
     */
    public String getUpdateUser();

    /**
     * Set the value of updateUser.
     * @param updateUser The value of updateUser. (NullAllowed)
     */
    public void setUpdateUser(String updateUser);

    /**
	 * Enable common column auto set up. {for after disable because the default is enabled}
	 */
    public void enableCommonColumnAutoSetup();

    /**
	 * Disable common column auto set up.
	 */
    public void disableCommonColumnAutoSetup();
	
    /**
	 * Can the entity set up common column by auto?
	 * @return The determination, true or false.
	 */
	public boolean canCommonColumnAutoSetup();
}
